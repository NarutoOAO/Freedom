import React, { useState } from 'react'
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import logo from '../../images/add-circle-outline.svg';
// define a component to search tutors
export default function TutorSearch(props) {
  const [show, setShow] = useState(false);
  const [selectedUser, setSelectedUser] = useState('');
  const courseNumber = props.courseNumber;
  const token = sessionStorage.getItem('token');
  const [users, setUsers] = useState(null);
// Open the modal
  const handleShow = async () => {
    setShow(true);
    console.log(courseNumber);
    // get the tutors
    const response = await fetch('http://127.0.0.1:5005/api/v1/tutor/'+ courseNumber, {
      method: 'GET',
      headers: {
        'Authorization': token,
      },
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      console.log(data);
      setUsers(data.data);
    }
  }
// close the modal
  const handleClose = () => {
    setUsers(null);
    setShow(false);
    setSelectedUser('');
  }
// define the function to select a tutor
  const handleSubmitUser = async () => {
    const apiUrl = 'http://localhost:5005/api/v1/group/'+props.id;
    try {
      const response = await fetch(apiUrl, {
        method: 'PUT',
        headers: {
          'Content-type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify({
          responsible_id: selectedUser.user_id,
          responsible_name: selectedUser.nick_name
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
    setSelectedUser('');
    handleClose();
    props.getGroupsFn();
  };
// define the function to select a tutor
  const handleUser =(user)=>{
    setSelectedUser(user);
  }

  return (
    <>
      <div style={{ display: 'flex', height: '40px'}}>
        <img src={logo} alt="add" style={{width:'25px'}}  onClick={handleShow}/>
      </div>
      <Modal show={show} onHide={handleClose} size='lg'>
        <Modal.Header closeButton>
          <Modal.Title>Tutor List</Modal.Title>
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
                <th scope="col">Id</th>
                <th scope="col">NickName</th>
                <th scope="col">Email</th>
                <th scope="col">Authority</th>
              </tr>
            </thead>
            <tbody>
              {users !== null &&
                users.map((user,index) => (
                  <tr key={index} onClick={() => handleUser(user)}
                  className={selectedUser.user_id === user.user_id ? 'selectedCourseEnroll' : ''}>
                    <th scope="row">{user.user_id}</th>
                    <td>{user.nick_name}</td>
                    <td>{user.email}</td>
                    <td>{user.authority}</td>
                  </tr>
                ))}
            </tbody>
          </table>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="primary" onClick={handleSubmitUser}>Submit</Button>
        </Modal.Footer>
      </Modal>
    </>
  )
}
