INSERT INTO LMSystem.comment (id, created_at, updated_at, deleted_at, post_id, author_id, content) VALUES (1, '2023-06-07 08:02:39', '2023-06-07 08:02:39', null, 1, 1, '可以');

INSERT INTO LMSystem.user (id, created_at, updated_at, deleted_at, email, password_digest, nick_name, authority, avatar) VALUES (1, '2023-06-07 07:57:10', '2023-06-07 07:57:10', null, '2371848072@qq.com', '$2a$12$by4PFggRoVb1pe6nev8fxuq8v9ilQDaHLmbzYcsY.VIqZ72OZ5VMG', 'Yihang Yu', 1, 'avatar.JPG');
INSERT INTO LMSystem.user (id, created_at, updated_at, deleted_at, email, password_digest, nick_name, authority, avatar) VALUES (2, '2023-06-08 00:38:01', '2023-06-08 02:57:05', null, '23718480721@qq.com', '$2a$12$pWrnlJW1fXhpwjypILytSOz4LWniU91rqNexf.s4KRJHAUEkp2JjG', 'Yihang', 0, 'user2/Yihang.jpg');
INSERT INTO LMSystem.user (id, created_at, updated_at, deleted_at, email, password_digest, nick_name, authority, avatar) VALUES (4, '2023-06-08 05:00:36', '2023-06-08 05:00:36', null, '1@qq.com', '$2a$12$mcDJt0TtZXoAxRklcfuke.s/acznq30Pvr58Ys9sbtSoJj88OWHG6', 'teacher1', 1, 'avatar.JPG');
INSERT INTO LMSystem.user (id, created_at, updated_at, deleted_at, email, password_digest, nick_name, authority, avatar) VALUES (5, '2023-06-08 05:00:46', '2023-06-08 05:00:46', null, '2@qq.com', '$2a$12$frkqD/eHLFfxX9h4bg2zRufZihgcC5SBdK.bRXF79DPgUAo79e6PW', 'teacher2', 1, 'avatar.JPG');
INSERT INTO LMSystem.user (id, created_at, updated_at, deleted_at, email, password_digest, nick_name, authority, avatar) VALUES (6, '2023-06-08 05:00:51', '2023-06-08 05:00:51', null, '3@qq.com', '$2a$12$hINcVefayFplpEaTTo7c8uN9uLAaqB8inHFNy6F54.51OGtfoULyO', 'teacher3', 1, 'avatar.JPG');
INSERT INTO LMSystem.user (id, created_at, updated_at, deleted_at, email, password_digest, nick_name, authority, avatar) VALUES (7, '2023-06-08 05:01:04', '2023-06-08 05:01:04', null, '11@qq.com', '$2a$12$PDUFCISRd8/l1R.tXk.Fv.yvp8eCY5nEfcNabBV.dYJwmsl48hiBa', 'student1', 0, 'avatar.JPG');
INSERT INTO LMSystem.user (id, created_at, updated_at, deleted_at, email, password_digest, nick_name, authority, avatar) VALUES (8, '2023-06-08 05:01:08', '2023-06-08 05:01:08', null, '22@qq.com', '$2a$12$asT6P3CT4f/p0DOQdBwqMuZ0hC2Hd2ptjrGDwKzG0l3sbEQVlqXkm', 'student2', 0, 'avatar.JPG');
INSERT INTO LMSystem.user (id, created_at, updated_at, deleted_at, email, password_digest, nick_name, authority, avatar) VALUES (9, '2023-06-08 05:01:16', '2023-06-08 05:01:16', null, '33@qq.com', '$2a$12$G/qPp3wpQTwF2BShlMjGMeqq8AJ9E6ZjILLoU8cndC4IRDAm9m1I6', 'student3', 0, 'avatar.JPG');

INSERT INTO LMSystem.course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img) VALUES (1, '2023-06-08 05:02:03', '2023-06-08 05:02:03', null, 9900, 'project', 4, 'teacher1', 'course.JPG');
INSERT INTO LMSystem.course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img) VALUES (2, '2023-06-08 05:02:24', '2023-06-08 05:02:24', null, 5836, 'AI', 4, 'teacher1', 'course.JPG');
INSERT INTO LMSystem.course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img) VALUES (3, '2023-06-08 05:02:38', '2023-06-08 05:02:38', null, 9517, 'Computer Vision', 4, 'teacher1', 'course.JPG');
INSERT INTO LMSystem.course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img) VALUES (4, '2023-06-08 05:03:13', '2023-06-08 05:03:13', null, 9334, 'Computer Network', 5, 'teacher2', 'course.JPG');
INSERT INTO LMSystem.course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img) VALUES (5, '2023-06-08 05:03:28', '2023-06-08 05:03:28', null, 9101, 'Algorithm', 5, 'teacher2', 'course.JPG');
INSERT INTO LMSystem.course (id, created_at, updated_at, deleted_at, course_number, course_name, teacher_id, teacher_name, course_img) VALUES (6, '2023-06-08 06:16:52', '2023-06-08 06:16:52', null, 9101, 'Algorithm', 5, 'teacher2', 'course.JPG');

INSERT INTO LMSystem.course_select (id, created_at, updated_at, deleted_at, course_number, course_name, course_img, teacher_id, teacher_name, student_id, status) VALUES (1, '2023-06-08 05:53:38', '2023-06-08 05:53:38', null, 9334, 'Computer Network', 'course.JPG', 5, 'teacher2', 8, 0);
INSERT INTO LMSystem.course_select (id, created_at, updated_at, deleted_at, course_number, course_name, course_img, teacher_id, teacher_name, student_id, status) VALUES (2, '2023-06-08 05:53:43', '2023-06-08 05:53:43', null, 5836, 'AI', 'course.JPG', 4, 'teacher1', 8, 0);
INSERT INTO LMSystem.course_select (id, created_at, updated_at, deleted_at, course_number, course_name, course_img, teacher_id, teacher_name, student_id, status) VALUES (3, '2023-06-08 05:53:47', '2023-06-08 05:53:47', null, 9900, 'project', 'course.JPG', 4, 'teacher1', 8, 0);

INSERT INTO LMSystem.forum (id, created_at, updated_at, deleted_at, forum_name, introduction) VALUES (1, '2023-06-07 08:02:34', '2023-06-07 08:02:34', null, 'quiz', 'discussion about uiz');

INSERT INTO LMSystem.post (id, created_at, updated_at, deleted_at, forum_id, title, content, author_id, status) VALUES (1, '2023-06-07 08:02:44', '2023-06-07 08:02:44', null, 1, '学习日志2', '今天学习了2', 1, 1);