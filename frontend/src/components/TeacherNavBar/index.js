import React from 'react';
import './style.css';
import {Link} from 'react-router-dom'
function TeacherNavBar(props) {
  const chooseMenu = (option) => {
    props.setOptionFn(option);
  };

  return (
    <div className="navigation_teacher">
      <ul>
        <li className='courseLabel'>Course ID: {props.courseNumber}</li>
        <li onClick={() => chooseMenu('OnlineClass')}>Tutors</li>
        <li onClick={() => chooseMenu('SetMaterial')}>Material</li>
        <li onClick={() => chooseMenu('EnterQuiz')}>Quiz</li>
        <li onClick={() => chooseMenu('SetAssignment')}>Assignment</li>
        <li onClick={() => chooseMenu('GiveMark')}>Mark</li>
        <li><Link to={`/discussion_forum/${props.courseNumber}`} className="link">Discussion Forum</Link></li>
      </ul>
    </div>
  );
}

export default TeacherNavBar;
