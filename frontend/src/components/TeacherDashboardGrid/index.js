import { useState, useEffect } from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { Link } from 'react-router-dom';
import './style.css';
// define the dashboard for teacher
function TeacherDashboardGrid() {
  const [courses, setCourses] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const coursesPerPage = 6; 
  const token = sessionStorage.getItem('token');
// get the courses
  useEffect(() => {
    const apiUrl = 'http://127.0.0.1:5005/api/v1/course';
    fetch(apiUrl, {
      method: 'GET',
      headers: {
        'Authorization': token
      }
    })
      .then(response => response.json())
      .then(data => {
        if (data.status === 200) {
          if(data.data!==null){
            setCourses(data.data);
          }
          // console.log(data);
        } else {
          alert(data.msg);
        }
      })
      .catch(error => {
        alert("Failed to fetch course data!");
      });
  }, []);
// define the pagination
  const indexOfLastCourse = currentPage * coursesPerPage;
  const indexOfFirstCourse = indexOfLastCourse - coursesPerPage;
  const currentCourses = courses.slice(indexOfFirstCourse, indexOfLastCourse);
  const totalPages = Math.ceil(courses.length / coursesPerPage);
// define the function to render the course
  const renderCourseRow = (course) => {
    return (
      <Col md={4} key={course.CourseNumber}>
        <Link to={`/teacher/course/${course.CourseNumber}`}>
          <div className="courseContainer">
            <img src={course.CourseImg} alt="course" className='courseImg' />
            <div className="courseInfo">
              <div className="courseDisplay">Course number:  {course.CourseNumber}</div>
              <div className="courseDisplay">Course name:  {course.CourseName}</div>
              <div className="courseDisplay">Teacher name:  {course.TeacherName}</div>
            </div>
          </div>
        </Link>
      </Col>
    );
  };
// define the function to change the page
  const handlePageChange = (pageNumber) => {
    setCurrentPage(pageNumber);
  };
// define the function to render the pagination
  const renderPagination = () => {
    const pageNumbers = [];
    for (let i = 1; i <= totalPages; i++) {
      pageNumbers.push(
        <li
          key={i}
          className={`paginationItem ${i === currentPage ? 'active' : ''}`}
          onClick={() => handlePageChange(i)}
        >
          {i}
        </li>
      );
    }
    return (
      <ul className="pagination">
        {pageNumbers}
      </ul>
    );
  };
// define the function to render the course
  const renderCourses = () => {
    const rows = [];
    for (let i = 0; i < currentCourses.length; i += 3) {
      const rowCourses = currentCourses.slice(i, i + 3);
      const row = (
        <Row key={`row-${i / 3}`}>
          {rowCourses.map(course => renderCourseRow(course))}
        </Row>
      );
      rows.push(row);
    }
    return rows;
  };

  return (
    <Container className='TeacherDashboardGrid'>
      {courses.length > 0 ? (
        <>
          {renderCourses()}
          {renderPagination()}
        </>
      ) : (
        <p>No courses available.</p>
      )}
    </Container>
  );
}

export default TeacherDashboardGrid;


