import React from 'react'
import { useParams } from 'react-router-dom';
import StudentNavBar from '../../../components/StudentNavBar'
import QuizCreate from '../../QuizCreate'
import OnlineClass from '../../OnlineClass'
import SetMaterial from '../../SetMaterial'
import SetAssignment from '../../SetAssignment'
import GiveMark from '../../GiveMark'
import './style.css'
export default function StudentCourse() {
  const [option,setOption] = React.useState('');
  const { courseNumber } = useParams();
  // console.log(courseNumber);
  React.useEffect(() => {
    document.getElementById('QuizCreate').style.display = 'none';
    document.getElementById('OnlineClass').style.display = 'none';
    document.getElementById('SetMaterial').style.display = 'none';
    document.getElementById('SetAssignment').style.display = 'none';
    document.getElementById('GiveMark').style.display = 'none';
    // console.log(document.getElementById(option).style.display);
    if(option!==''){
      // console.log(option);
      document.getElementById(option).style.display = 'block';
    }

  }, [option]);
  return (
    <div className='teacherCourseAll'>
      <StudentNavBar setOptionFn={setOption} courseNumber={courseNumber}/>
      <div id='QuizCreate' className='teacherCourseContent' style={{ display: 'none' }}>
        <QuizCreate/>
      </div>
      <div id='OnlineClass' className='teacherCourseContent' style={{ display: 'none' }}>
        <OnlineClass/>
      </div>
      <div id='SetMaterial' className='teacherCourseContent' style={{ display: 'none' }}>
        <SetMaterial/>
      </div>
      <div id='SetAssignment' className='teacherCourseContent' style={{ display: 'none' }}>
        <SetAssignment/>
      </div>
      <div id='GiveMark' className='teacherCourseContent' style={{ display: 'none' }}>
        <GiveMark/>
      </div>
    </div>
  )
}
