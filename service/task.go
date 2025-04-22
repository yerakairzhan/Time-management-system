package service

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"TimeManagementSystem/repository"
	"fmt"
	"time"
)

type TaskServiceImpl struct {
	taskRepo repository.TaskRepository
	userRepo repository.UserRepository
}

func NewTaskService(taskRepo repository.TaskRepository, userRepo repository.UserRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

func (s *TaskServiceImpl) CreateTask(userId int, task sqlc.Task) (int, error) {
	return s.taskRepo.Create((userId), task)
}

func (s *TaskServiceImpl) GetTasksByUserID(userId int) ([]sqlc.Task, error) {
	return s.taskRepo.GetTasksByUserID(userId)
}

func (s *TaskServiceImpl) GetTaskById(taskID int) (sqlc.Task, error) {
	return s.taskRepo.GetTaskById(taskID)
}

func (s *TaskServiceImpl) UpdateTask(taskId int, task sqlc.Task) error {
	return s.taskRepo.Update(taskId, task)
}

func (s *TaskServiceImpl) DeleteTask(taskId int) error {
	return s.taskRepo.Delete(taskId)
}

func (s *TaskServiceImpl) StartTask(taskId int) error {
	activeTimer, err := s.taskRepo.GetActiveTimer(taskId)
	if err != nil {
		return fmt.Errorf("error checking active timer: %w", err)
	}

	if activeTimer.ID != 0 {
		return fmt.Errorf("task timer is already running")
	}

	err = s.taskRepo.StartTimer(taskId)
	if err != nil {
		return fmt.Errorf("error starting task timer: %w", err)
	}

	return nil
}

func (s *TaskServiceImpl) StopTask(taskId int) error {
	activeTimer, err := s.taskRepo.GetActiveTimer(taskId)
	if err != nil {
		return fmt.Errorf("error checking active timer: %w", err)
	}

	if activeTimer.ID == 0 {
		return fmt.Errorf("no active timer found for this task")
	}

	err = s.taskRepo.StopTimer(taskId)
	if err != nil {
		return fmt.Errorf("error stopping task timer: %w", err)
	}

	return nil
}

func (s *TaskServiceImpl) GetTimeSpent(taskId int) ([]time.Duration, error) {
	timeLogs, err := s.taskRepo.GetTimeSpent(taskId)
	if err != nil {
		return nil, err
	}

	var durations []time.Duration
	for _, log := range timeLogs {
		if !log.StartTime.Valid || !log.EndTime.Valid {
			continue
		}
		duration := log.EndTime.Time.Sub(log.StartTime.Time)
		durations = append(durations, duration)
	}

	if len(durations) == 0 {
		return nil, fmt.Errorf("no valid time logs found")
	}

	return durations, nil
}
