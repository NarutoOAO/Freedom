import React, { useState } from 'react'
import logo from '../../images/person-add-outline.svg'
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
export default function UserSearch(props) {
  const [search, setSearch] = useState('');
  const [show, setShow] = useState(false);
  const [selectedUser, setSelectedUser] = useState('');
  const courseNumber = props.courseNumber;
  const token = sessionStorage.getItem('token');
  const [users, setUsers] = useState(null);

  const handleShow = async () => {
    console.log(search);
    const response = await fetch('http://127.0.0.1:5005/api/v1/users', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body: JSON.stringify({
        info: search
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      setUsers(data.data.item);
      setShow(true);
    }
  }

  const handleClose = () => {
    setSearch('');
    setUsers(null);
    setShow(false);
    setSelectedUser('');
  }

  const handleSubmitUser = async () => {
    const apiUrl = 'http://localhost:5005/api/v1/tutor';
    try {
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify({
          user_id:selectedUser.id,
          email:selectedUser.email,
          nick_name: selectedUser.nickname,
          authority: selectedUser.authority,
          course_number: parseInt(courseNumber),
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
    handleShow();
    props.getTutorsFn();
  };

  const handleUser =(user)=>{
    setSelectedUser(user);
  }

  return (
    <>
      <div style={{ display: 'flex', height: '40px', marginBottom:'20px'}}>
        <img src={logo} alt="search logo" style={{  marginRight: '10px', width:'30px' }} onClick={handleShow} />
        <input type='search' placeholder='Invite tutor' style={{  width: '20%'}} onChange={(event) => setSearch(event.target.value)} value={search} />
      </div>
      <Modal show={show} onHide={handleClose} size='lg'>
        <Modal.Header closeButton>
          <Modal.Title>User List</Modal.Title>
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
                  className={selectedUser.id === user.id ? 'selectedCourseEnroll' : ''}>
                    <th scope="row">{user.id}</th>
                    <td>{user.nickname}</td>
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
