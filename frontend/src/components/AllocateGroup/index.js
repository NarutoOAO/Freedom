import React, { useState } from 'react'
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import logo from '../../images/add-circle-outline.svg';
export default function AllocateGroup(props) {
  const [show, setShow] = useState(false);
  const [selectedGroup, setSelectedGroup] = useState('');
  const courseNumber = props.courseNumber;
  const token = sessionStorage.getItem('token');
  const [groups, setGroups] = useState(null);
  const handleShow = async () => {
    setShow(true);
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
      console.log(data);
      setGroups(data.data);
    }
  }

  const handleClose = () => {
    setGroups(null);
    setShow(false);
    setSelectedGroup('');
  }

  const handleSubmit = async () => {
    const apiUrl = 'http://localhost:5005/api/v1/assignment_solution';
    try {
      const response = await fetch(apiUrl, {
        method: 'PUT',
        headers: {
          'Content-type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify({
          ass_mark_id:props.id,
          group_id:selectedGroup.id,
          group_name: selectedGroup.group_name
        })
      });
      const data = await response.json();
      if (data.status===200) {
        alert('Succeed!');
      } else {
        alert(data.msg);
      }
    } catch (error) {
      alert('Failed. Please try again.');
      console.error(error);
    }
    setSelectedGroup('');
    handleClose();
    props.mark_assigmentFn();
  };

  const handleGroup =(group)=>{
    setSelectedGroup(group);
  }

  return (
    <>
      <div style={{ display: 'flex', height: '40px'}}>
        <img src={logo} alt="add" style={{width:'25px'}}  onClick={handleShow}/>
      </div>
      <Modal show={show} onHide={handleClose} size='lg'>
        <Modal.Header closeButton>
          <Modal.Title>Group List</Modal.Title>
        </Modal.Header>
        <Modal.Body className='modalBody' style={{ maxHeight: '500px', overflowY: 'auto' }}>
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
                    <tr key={index} onClick={() => handleGroup(group)}
                    className={selectedGroup.id === group.id ? 'selectedCourseEnroll' : ''}>
                      <th scope="row">{group.id}</th>
                      <td>{group.group_name}</td>
                      <td>{group.teacher_name}</td>
                      <td>{group.responsible_name}</td>
                    </tr>
                  ))}
              </tbody>
            </table>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="primary" onClick={handleSubmit}>Submit</Button>
        </Modal.Footer>
      </Modal>
    </>
  )
}
