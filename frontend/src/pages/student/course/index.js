import React from 'react'
import { useParams } from 'react-router-dom';
import StudentNavBar from '../../../components/StudentNavBar'
import EnterQuizStudent from '../../EnterQuizStudent'
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
    document.getElementById('EnterQuizStudent').style.display = 'none';
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
      <div id='EnterQuizStudent' className='teacherCourseContent' style={{ display: 'none' }}>
        <EnterQuizStudent/>
      </div>
      <div id='OnlineClass' className='teacherCourseContent' style={{ display: 'none' }}>
        <OnlineClass/>
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
