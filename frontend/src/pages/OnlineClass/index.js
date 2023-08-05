import React, { useEffect, useState } from 'react'
import UserSearch from '../../components/SearchUser';
import CreateGroup from '../../components/CreateGroup';
import Accordion from 'react-bootstrap/Accordion';
import TutorSearch from '../../components/SearchTutor';
// define a page for teacher to manage tutor and group
export default function OnlineClass(props) {
  const token = sessionStorage.getItem('token');
  const [course, setCourse] = useState('');
  const courseNumber = props.courseNumber;
  const user_id = sessionStorage.getItem('user_id');
  const [tutors, setTutors] = useState(null);
  const [groups, setGroups] = useState(null);
  // hook for update
  useEffect(()=>{
    getCourse();
    getTutors();
    getGroups();
  },[props.option]);
// fetch course information from the server
  const getCourse = async()=>{
    const res = await fetch("http://localhost:5005/api/v1/course/"+courseNumber,{
      method: 'GET',
      headers: {
        'Authorization': token
      }
    });
    const data = await res.json();
    if(data.status!==200){
      alert(data.msg);
    }else{
      setCourse(data.data);
    }
  }
// fetch tutors information from the server
  const getTutors = async()=>{
    const res = await fetch("http://localhost:5005/api/v1/tutor/"+courseNumber,{
      method: 'GET',
      headers: {
        'Authorization': token
      }
    });
    const data = await res.json();
    if(data.status!==200){
      alert(data.msg);
    }else{
      setTutors(data.data);
    }
  }
// fetch groups information from the server
  const getGroups = async()=>{
    const res = await fetch("http://localhost:5005/api/v1/group/"+courseNumber,{
      method: 'GET',
      headers: {
        'Authorization': token
      }
    });
    const data = await res.json();
    if(data.status!==200){
      alert(data.msg);
    }else{
      setGroups(data.data);
    }
  }
// render the page
  const renderAddorTutor=(group)=>{
    if (group.responsible_name===""){
      return <span><TutorSearch courseNumber={courseNumber} id={group.id}   getGroupsFn={getGroups}/></span>
    }else{
      return <span>{group.responsible_name}</span>
    }
  }

  return (
    <>
    {parseInt(user_id) === course.TeacherId &&
    <Accordion defaultActiveKey="0" style={{marginTop:"20px", width:'95%'}}>
      <Accordion.Item eventKey="0">
        <Accordion.Header>Tutors</Accordion.Header>
        <Accordion.Body>
          <UserSearch courseNumber={courseNumber} getTutorsFn={getTutors} />
        <table className="table table-hover">
        <colgroup>
          <col style={{ width: '10%' }} />
          <col style={{ width: '20%' }} />
          <col style={{ width: '30%' }} />
          <col style={{ width: '20%' }} />
        </colgroup>
        <thead>
          <tr>
            <th scope="col">User Id</th>
            <th scope="col">NickName</th>
            <th scope="col">Email</th>
            <th scope="col">Authority</th>
          </tr>
        </thead>
        <tbody>
          {tutors!==null && tutors!==[] && tutors!== undefined && tutors!==''&&
            tutors.map((tutor, index) => (
              <tr key={index}>
                <th scope="row">{tutor.user_id}</th>
                <td>{tutor.nick_name}</td>
                <td>{tutor.email}</td>
                <td>{tutor.authority===1?'teacher':'student'}</td>
              </tr>
            ))}
        </tbody>
      </table>
        </Accordion.Body>
      </Accordion.Item>
      <Accordion.Item eventKey="1">
        <Accordion.Header>Groups</Accordion.Header>
        <Accordion.Body>
          <CreateGroup courseNumber={courseNumber}  getGroupsFn={getGroups}/>
      <table className="table table-hover">
        <colgroup>
          <col style={{ width: '10%' }} />
          <col style={{ width: '20%' }} />
          <col style={{ width: '30%' }} />
          <col style={{ width: '20%' }} />
        </colgroup>
        <thead>
          <tr>
            <th scope="col">Group Id</th>
            <th scope="col">Group Name</th>
            <th scope="col">Teacher</th>
            <th scope="col">Tutor</th>
          </tr>
        </thead>
        <tbody>
          {groups!==null && groups!==[] && groups!== undefined && groups!==''&&
            groups.map((group, index) => (
              <tr key={index}>
                <th scope="row">{group.id}</th>
                <td>{group.group_name}</td>
                <td>{group.teacher_name}</td>
                <td>{renderAddorTutor(group)}</td>
              </tr>
            ))}
        </tbody>
      </table>
        </Accordion.Body>
      </Accordion.Item>
    </Accordion>}
    </>
  )
}