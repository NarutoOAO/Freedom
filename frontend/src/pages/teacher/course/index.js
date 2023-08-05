import React from 'react'
import { useParams } from 'react-router-dom';
import TeacherNavBar from '../../../components/TeacherNavBar'
import EnterQuiz from '../../EnterQuiz'
import OnlineClass from '../../OnlineClass'
import SetMaterial from '../../SetMaterial'
import SetAssignment from '../../SetAssignment'
import GiveMark from '../../GiveMark'
import './style.css'
// define the course page for teacher
export default function TeacherCourse() {
  const [option,setOption] = React.useState('');
  const { courseNumber } = useParams();
// hook to set the option
  React.useEffect(() => {
    document.getElementById('EnterQuiz').style.display = 'none';
    document.getElementById('OnlineClass').style.display = 'none';
    document.getElementById('SetMaterial').style.display = 'none';
    document.getElementById('SetAssignment').style.display = 'none';
    document.getElementById('GiveMark').style.display = 'none';
    // console.log(document.getElementById(option).style.display);
    if(option!==''){
      document.getElementById(option).style.display = 'block';
    }

  }, [option]);
  return (
    <div className='teacherCourseAll'>
      <TeacherNavBar setOptionFn={setOption} courseNumber={courseNumber}/>
      <div id='EnterQuiz' className='teacherCourseContent' style={{ display: 'none' }}>
        <EnterQuiz/>
      </div>
      <div id='OnlineClass' className='teacherCourseContent' style={{ display: 'none' }}>
        <OnlineClass courseNumber={courseNumber} option={option}/>
      </div>
      <div id='SetMaterial' className='teacherCourseContent' style={{ display: 'none' }}>
        <SetMaterial option={option}/>
      </div>
      <div id='SetAssignment' className='teacherCourseContent' style={{ display: 'none' }}>
        <SetAssignment option={option}/>
      </div>
      <div id='GiveMark' className='teacherCourseContent' style={{ display: 'none' }}>
        <GiveMark courseNumber={courseNumber} option={option}/>
      </div>
    </div>
  )
}
