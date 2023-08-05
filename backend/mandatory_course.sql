create table mandatory_course
(
    id              bigint unsigned auto_increment
        primary key,
    created_at      datetime        null,
    updated_at      datetime        null,
    deleted_at      datetime        null,
    course_number   bigint          null,
    course_name     varchar(256)    null,
    teacher_id      bigint unsigned null,
    teacher_name    varchar(256)    null,
    course_img      varchar(256)    null,
    class_time      varchar(256)    null,
    course_location varchar(256)    null,
    classification  varchar(256)    null,
    max_people      bigint          null,
    term            varchar(256)    null,
    constraint course_number
        unique (course_number)
);

create index idx_mandatory_course_deleted_at
    on mandatory_course (deleted_at);

INSERT INTO LMSystem.mandatory_course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img, class_time, course_location, classification, max_people, term) VALUES (2, '2023-07-19 04:51:18', '2023-07-19 04:51:18', null, 9021, 'Principles of Programming', 1, '19', 'course.JPG', 'Monday: 8:00 - 10:00
Wednesday: 10:00 - 12:00', 'Science Th', 'AI,IT', 100, '1,2,3');
INSERT INTO LMSystem.mandatory_course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img, class_time, course_location, classification, max_people, term) VALUES (3, '2023-07-19 04:51:18', '2023-07-19 04:51:18', null, 9024, 'Data Structure and Algorithm', 1, '19', 'course.JPG', 'Monday: 8:00 - 10:00
Wednesday: 10:00 - 12:00', 'Science Th', 'AI,IT', 300, '2');
INSERT INTO LMSystem.mandatory_course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img, class_time, course_location, classification, max_people, term) VALUES (4, '2023-07-30 01:01:00', '2023-07-30 01:01:02', null, 9032, 'Microprocessors and Interfacing', 1, '19', 'course.JPG', 'Monday: 8:00 - 10:00
Wednesday: 10:00 - 12:00', 'Science Th', 'AI,IT', 200, '3');
INSERT INTO LMSystem.mandatory_course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img, class_time, course_location, classification, max_people, term) VALUES (5, '2023-07-30 15:26:30', '2023-07-30 15:26:33', null, 9311, 'Database System', 1, '19', 'course.JPG', 'Monday: 10:00 - 12:00', 'Col LG02', 'AI,IT', 300, '1,2,3');
INSERT INTO LMSystem.mandatory_course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img, class_time, course_location, classification, max_people, term) VALUES (6, '2023-07-30 15:32:00', '2023-07-30 15:32:02', null, 9417, 'Machine  Learning and Data Mining', 1, '19', 'course.JPG', 'Wednesday: 10:00 - 12:00', 'Griff M18', 'AI', 280, '2,3');
INSERT INTO LMSystem.mandatory_course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img, class_time, course_location, classification, max_people, term) VALUES (7, '2023-07-30 15:32:56', '2023-07-30 15:32:58', null, 9900, 'Information Technology Project', 1, '19', 'course.JPG', 'Tuesday: 12:00 - 14:00
Thursday: 14:00 - 16:00', 'Col LG02', 'AI,IT', 340, '2,3');
INSERT INTO LMSystem.mandatory_course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img, class_time, course_location, classification, max_people, term) VALUES (8, '2023-07-30 15:33:27', '2023-07-30 15:33:29', null, 9414, 'Artificial Intelligence', 1, '19', 'course.JPG', 'Wednesday: 10:00 - 12:00
Thursday: 14:00 - 16:00', 'Griff M18', 'AI', 270, '1,2,3');
INSERT INTO LMSystem.mandatory_course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img, class_time, course_location, classification, max_people, term) VALUES (9, '2023-07-30 15:33:31', '2023-07-30 15:33:33', null, 9020, 'Foundations of Computer Science', 1, '19', 'course.JPG', 'Tuesday: 10:00 - 12:00
Friday: 14:00 - 16:00', 'Science Th', 'IT', 300, '1,2');
