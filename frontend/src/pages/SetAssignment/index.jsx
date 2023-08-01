
import AssigmentInformation from '../../components/AssigmentInformation';
import "./setAssigment.css";
import PostAssigment from '../../components/PostAssigment';
import React, { useState } from 'react';
import { useParams } from 'react-router-dom';

export default function SetAssigment (props) {
  const authority=sessionStorage.getItem('authority');
  const { courseNumber } = useParams();
  const [assigmentFlag, setAssigmentFlag] = useState(0);


  return (
    <div className = "assigment-container">
      {authority !== '0' && (
      <div className="postAssigment-container">
            <PostAssigment courseNumber={courseNumber} setAssigmentFlag={setAssigmentFlag} assigmentFlag={assigmentFlag} option={props.option}/>
      </div>
      )}
      {/* <input type="datetime-local" id="myDateTimePicker"></input> */}
      <div className="assigmentInforamtion-container">
            <AssigmentInformation courseNumber={courseNumber} setAssigmentFlag={setAssigmentFlag} assigmentFlag={assigmentFlag} option={props.option}/>
      </div>
    </div>
  )
}
