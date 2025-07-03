package assignment

import (
	"context"
	"fmt"
	"sync"
	"time"

	assignmentDom "github.com/NupalHariz/DD/src/business/domain/assignment"
	assignmentCategoryDom "github.com/NupalHariz/DD/src/business/domain/assignment_category"
	userDom "github.com/NupalHariz/DD/src/business/domain/user"
	"github.com/NupalHariz/DD/src/business/dto"
	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/NupalHariz/DD/src/business/service/mail"
	"github.com/reyhanmichiels/go-pkg/v2/auth"
	"github.com/reyhanmichiels/go-pkg/v2/query"
)

type Interface interface {
	Create(ctx context.Context, param dto.CreateAssignmentParam) error
	Update(ctx context.Context, param dto.UpdateAssignmentParam) error
	GetAll(ctx context.Context, param dto.GetAllAssignmentParam) ([]dto.GetAllAssignmentResponse, error)
	TodayDeadlineScheduler(ctx context.Context) error
}

type assignment struct {
	auth                  auth.Interface
	assignmentDom         assignmentDom.Interface
	assignmentCategoryDom assignmentCategoryDom.Interface
	userDom               userDom.Interface
	mail                  mail.Interface
}

type InitParam struct {
	Auth                  auth.Interface
	Assignment            assignmentDom.Interface
	AssignmentCategoryDom assignmentCategoryDom.Interface
	UserDom               userDom.Interface
	Mail                  mail.Interface
}

func Init(param InitParam) Interface {
	return &assignment{
		auth:                  param.Auth,
		assignmentDom:         param.Assignment,
		assignmentCategoryDom: param.AssignmentCategoryDom,
		userDom:               param.UserDom,
		mail:                  param.Mail,
	}
}

func (a *assignment) Create(ctx context.Context, param dto.CreateAssignmentParam) error {
	loginUser, err := a.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	assignmentInputParam := param.ToAssignmentInputParam(loginUser.ID)

	err = a.assignmentDom.Create(ctx, assignmentInputParam)
	if err != nil {
		return err
	}

	return nil
}

func (a *assignment) Update(ctx context.Context, param dto.UpdateAssignmentParam) error {
	assignmentUpdateParam := param.ToAssignmentUpdateParam()

	err := a.assignmentDom.Update(ctx, assignmentUpdateParam, entity.AssignmentParam{Id: param.Id})
	if err != nil {
		return err
	}

	return nil
}

func (a *assignment) GetAll(ctx context.Context, param dto.GetAllAssignmentParam) ([]dto.GetAllAssignmentResponse, error) {
	var datas []dto.GetAllAssignmentResponse

	loginUser, err := a.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return datas, err
	}

	assignments, err := a.assignmentDom.GetAll(
		ctx,
		entity.AssignmentParam{
			UserId: loginUser.ID,
			Option: query.Option{
				DisableLimit: false,
			},
			PaginationParam: param.PaginationParam,
		})
	if err != nil {
		return datas, err
	}

	categoryIdMapSet := make(map[int64]struct{})
	for _, a := range assignments {
		categoryIdMapSet[a.CategoryId] = struct{}{}
	}

	var categoryIds []int64
	for id := range categoryIdMapSet {
		categoryIds = append(categoryIds, id)
	}

	categories, err := a.assignmentCategoryDom.GetAll(ctx, entity.AssignmentCategoryParam{Ids: categoryIds})
	if err != nil {
		return datas, err
	}

	categoryMap := make(map[int64]string)
	for _, c := range categories {
		categoryMap[c.Id] = c.Name
	}

	for _, a := range assignments {
		categoryName := categoryMap[a.CategoryId]

		data := dto.GetAllAssignmentResponse{
			Id:       a.Id,
			Category: categoryName,
			Name:     a.Name,
			Deadline: a.Deadline,
			Status:   string(a.Status),
			Priority: string(a.Priority),
		}

		datas = append(datas, data)
	}

	return datas, nil
}

func (a *assignment) TodayDeadlineScheduler(ctx context.Context) error {
	currentDate := time.Now().Format("2006-01-02")

	assignments, err := a.assignmentDom.GetAll(ctx, entity.AssignmentParam{Deadline: currentDate, Status: string(entity.OnGoing)})
	if err != nil {
		return err
	}

	userAssignments := make(map[int64][]entity.Assignment)
	for _, a := range assignments {
		userAssignments[a.UserId] = append(userAssignments[a.UserId], a)
	}

	var userIds []int64
	for uA := range userAssignments {
		userIds = append(userIds, uA)
	}

	users, _, err := a.userDom.GetList(ctx, entity.UserParam{Ids: userIds})
	if err != nil {
		return err
	}

	userMap := make(map[int64]entity.User)
	for _, u := range users {
		userMap[u.ID] = u
	}

	var wg sync.WaitGroup

	for userId, assignment := range userAssignments {
		wg.Add(1)

		user := userMap[userId]

		body := fmt.Sprintf("Hello %s, you have the following assignment that due todays: <br>", user.Name)
		for _, a := range assignment {
			body += fmt.Sprintf("- %s<br>", a.Name)
		}

		go func() {
			defer wg.Done()
			_ = a.mail.SendEmail(ctx, user.Email, "Assignment Deadline", body)
		}()
	}

	wg.Wait()

	return nil
}
