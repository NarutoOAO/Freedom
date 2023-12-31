import React from 'react'
import './style.css';
import {Link} from 'react-router-dom'
// define the navigation bar for student
export default function StudentNavBar(props) {
  const chooseMenu = (option) => {
    props.setOptionFn(option);
  };

  return (
    <div className="navigation_teacher">
      <ul>
        <li className='courseLabel'> Course ID: {props.courseNumber}</li>
        <li onClick={() => chooseMenu('SetMaterial')}>Material</li>
        <li onClick={() => chooseMenu('EnterQuizStudent')}>Quiz</li>
        <li onClick={() => chooseMenu('SetAssignment')}>Assignment</li>
        <li onClick={() => chooseMenu('GiveMark')}>View marks</li>
        <li><Link to={`/discussion_forum/${props.courseNumber}`} className="link">Discussion Forum</Link></li>
      </ul>
    </div>
  );
}
