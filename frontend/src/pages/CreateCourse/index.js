import React, { useState } from 'react';
import './style.css'
function CreateCourse() {
  const [courseNumber, setCourseNumber] = useState('');
  const [courseName, setCourseName] = useState('');
  const token = sessionStorage.getItem('token');
  const [maxEnrollPeople,setMaxEnrollPeople]= useState('');
  const [courseLocation,setCourseLocation]= useState('');
  const [classification,setClassification]= useState('');
  const [classtime, setClasstime] = useState([{ day: '', time: '' }]);

  const addClassTime = () => {
    setClasstime([...classtime, { day: '', time: '' }]);
  };

  const removeClassTime = (index) => {
    if (classtime.length > 1) {
      const updatedClasstime = [...classtime];
      updatedClasstime.splice(index, 1);
      setClasstime(updatedClasstime);
    }
  };
  

  const handleCreateCourse = async () => {
    console.log(classtime)
    // Check empty for class times
    const hasEmptyDayOrTime = classtime.some((time) => !time.day || !time.time);

    if (hasEmptyDayOrTime) {
      alert('Every day and time in Class Time must be selected.');
      return;
    }
    // Check for duplicate class times
    const hasDuplicates = classtime.some(
      (time, index) =>
        classtime.findIndex(
          (otherTime, otherIndex) =>
            index !== otherIndex &&
            time.day === otherTime.day &&
            time.time === otherTime.time
        ) !== -1
    );
    if (hasDuplicates) {
      alert('Duplicate class times are not allowed.');
      return;
    }
    // Sort class times
    const sortedClasstime = classtime.slice().sort((a, b) => {
      // Compare days first
      const dayOrder = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'];
      const dayA = dayOrder.indexOf(a.day);
      const dayB = dayOrder.indexOf(b.day);
      if (dayA !== dayB) {
        return dayA - dayB;
      }
  
      // Days are same, compare times
      const timeOrder = ['8:00 - 10:00', '10:00 - 12:00','12:00 - 14:00','16:00 - 18:00','18:00 - 20:00'];
      const timeA = timeOrder.indexOf(a.time);
      const timeB = timeOrder.indexOf(b.time);
      return timeA - timeB;
    });
   
   
    const classTimeString = sortedClasstime.map((time) => `${time.day}: ${time.time}`).join('\n');
    console.log(sortedClasstime)
    console.log(typeof(sortedClasstime))
    if (parseInt(courseNumber) === 0 || isNaN(parseInt(courseNumber))) {
      alert('Please enter a valid course number');
      return;
    }

    if (!courseLocation) {
      alert('Course Location cannot be empty.');
      return;
    }

    if (!classification) {
      alert('Classification cannot be empty.');
      return;
    }

    if (parseInt(maxEnrollPeople) < 1 || isNaN(parseInt(maxEnrollPeople))) {
      alert('Max Enroll People should be a valid integer greater than or equal to 1.');
      return;
    }
    const response = await fetch('http://127.0.0.1:5005/api/v1/teacher-course', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body: JSON.stringify({
        course_number: parseInt(courseNumber),
        course_name: courseName,
    	  class_time : classTimeString,
        course_location: courseLocation,
        max_people: parseInt(maxEnrollPeople),
        Classification: classification,
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      // console.log(data);
      alert("Succeed!");
    }
    setCourseName('');
    setCourseNumber('');
    setMaxEnrollPeople('');
    setCourseLocation('');
    setClassification('');
    setClasstime([{ day: '', time: '' }]);
  };

  return (
    <div className='CreateCourse'>
      <h2>Create course</h2>
      <form>
        <div> 
          <label>
            Course Number:
            <input
              type="text"
              value={courseNumber}
              onChange={(e) => setCourseNumber(e.target.value)}
            />
        </label>
        </div>
        <div>
          <label>
          Course Name:
            <input
              type="text"
              value={courseName}
              onChange={(e) => setCourseName(e.target.value)}
            />
          </label>
        </div>
        <div>
          <label>
          Max Enroll People:
            <input
              type="text"
              value={maxEnrollPeople}
              onChange={(e) => setMaxEnrollPeople(e.target.value)}
            />
          </label>
        </div>
        <div>
          <label>
          Course Location:
              <select id="locationSelect" value={courseLocation} onChange={(e) => setCourseLocation(e.target.value)} style={{width:'100%', fontSize:'17pt', marginBottom:'10px', fontWeight:'normal'}}>
                  <option value="">Select Location</option>
                  <option value="Griff M18">Griff M18</option>
                  <option value="Col LG02">Col LG02</option>
                  <option value="Science Th">Science Th</option>
                  <option value="Physics Th">Physics Th</option>
                  <option value="BUS 105">BUS 105</option>
              </select>
          </label>
        </div>
        <div>
          <label>
          Classification:
              <select id="classificationSelect" value={classification} onChange={(e) => setClassification(e.target.value)} style={{width:'100%', fontSize:'17pt', marginBottom:'10px', fontWeight:'normal'}}>
                  <option value="">Select Classification</option>
                  <option value="ADK">ADK</option>
                  <option value="DE">DE</option>
                  <option value="Core Courses">Core Courses</option>
              </select>
          </label>
        </div>
        {classtime.map((time, index) => (
          <div key={index}>
            <label>
              Class Day:
              <select
                value={time.day}
                onChange={(e) => {
                  const updatedClasstime = [...classtime];
                  updatedClasstime[index].day = e.target.value;
                  setClasstime(updatedClasstime);
                }}
              >
                <option value=''>Select Day</option>
                <option value='Monday'>Monday</option>
                <option value='Tuesday'>Tuesday</option>
                <option value='Wednesday'>Wednesday</option>
                <option value='Thursday'>Thursday</option>
                <option value='Friday'>Friday</option>
              </select>
            </label>
            <label>
              Class Time:
              <select
                value={time.time}
                onChange={(e) => {
                  const updatedClasstime = [...classtime];
                  updatedClasstime[index].time = e.target.value;
                  setClasstime(updatedClasstime);
                }}
              >
                <option value=''>Select Time</option>
                <option value='8:00 - 10:00'>8:00 - 10:00</option>
                <option value='10:00 - 12:00'>10:00 - 12:00</option>
                <option value='12:00 - 14:00'>12:00 - 14:00</option>
                <option value='14:00 - 16:00'>14:00 - 16:00</option>
                <option value='16:00 - 18:00'>16:00 - 18:00</option>
                <option value='18:00 - 20:00'>18:00 - 20:00</option>
              </select>
            </label>
            {index > 0 && (
              <button type='button' style={{marginLeft:'10px',marginTop:'10px'}} onClick={() => removeClassTime(index)}>
                Remove
              </button>
            )}
          </div>
        ))}

      
        <div style={{marginBottom:'10px',marginTop:'10px'}}>
          <button type='button' onClick={addClassTime}>
            Add Class Time
          </button>
        </div>
        <button type="button" onClick={handleCreateCourse}>
          Create
        </button>
      </form>
    </div>
  );
}

export default CreateCourse;
