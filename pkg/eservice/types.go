package eservice

type Group struct {
    Id          int    `json:"id"`
    Name        string `json:"name"`
    GroupRealId int    `json:"group_real_id"`
}

type Lesson struct {
    Id          int    `json:"id"`
    Day         string `json:"day"`
    Week        string `json:"week"`
    Time        int `json:"time"`
    Faculty     string `json:"faculty"`
    Lesson      string `json:"lesson"`
    TypeWork    string `json:"type_work"`
    LessonId    int    `json:"lesson_id"`
    Teacher     string `json:"teacher"`
    TeacherId   int    `json:"teacher_id"`
    Group       string `json:"group"`
    GroupId     int    `json:"group_id"`
    AuditCorps  string `json:"auditCorps"`
    AuditoryId  int    `json:"auditory_id"`
    PublishDate string `json:"publishDate"`
}

type ScheduleDay struct {
    Day     string   `json:"day"`
    Lessons []Lesson `json:"lessons"`
}
