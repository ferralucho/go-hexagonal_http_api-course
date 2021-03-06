package mooc

import "context"

// Course is the data structure that represents a course.
type Course struct {
	id       string
	name     string
	duration string
}

// NewCourse creates a new course.
func NewCourse(id, name, duration string) Course {
	return Course{
		id:       id,
		name:     name,
		duration: duration,
	}
}

// CourseRepository defines the expected behaviour from a course storage.
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

// ID returns the course unique identifier.
func (c Course) ID() string {
	return c.id
}

// Name returns the course name.
func (c Course) Name() string {
	return c.name
}

// Duration returns the course duration.
func (c Course) Duration() string {
	return c.duration
}
