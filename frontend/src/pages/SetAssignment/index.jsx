
import AssigmentInformation from '../../components/AssigmentInformation';
import "./setAssigment.css";
import PostAssigment from '../../components/PostAssigment';
import React, { useState } from 'react';
import { useParams } from 'react-router-dom';

export default function SetAssigment (props) {
  // Get the authority from session storage 
  const authority=sessionStorage.getItem('authority');
  const { courseNumber } = useParams();
  // Flag use for refresh page when assigment inforamtion change
  const [assigmentFlag, setAssigmentFlag] = useState(0);


  return (
    <div className = "assigment-container">
      {/* As teachers, they can post assigment*/}
      {authority !== '0' && (
      <div className="postAssigment-container">
            <PostAssigment courseNumber={courseNumber} setAssigmentFlag={setAssigmentFlag} assigmentFlag={assigmentFlag} option={props.option}/>
      </div>
      )}
       {/*See assigment inforamtion*/}
      <div className="assigmentInforamtion-container">
            <AssigmentInformation courseNumber={courseNumber} setAssigmentFlag={setAssigmentFlag} assigmentFlag={assigmentFlag} option={props.option}/>
      </div>
    </div>
  )
}
