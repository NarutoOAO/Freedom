import { useState } from 'react';
import Offcanvas from 'react-bootstrap/Offcanvas';
import logo from '../../images/notifications.svg'
import Toast from 'react-bootstrap/Toast';
function OffCanvasExample({ name, ...props }) {
  const [show, setShow] = useState(false);
  const [notifications, setNotifications] = useState('');
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);
  const token = sessionStorage.getItem('token');
  const getNotifications = async() =>{
    handleShow();
    const res = await fetch('http://127.0.0.1:5005/api/v1/notification', {
      method: 'GET',
      headers: {
        'Authorization': token,
      },
    })
    const data = await res.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      console.log(data.data);
      if (data.data!==null){
        setNotifications(data.data.reverse());
      }
    }
  }
  const renderStatus =(status)=>{
    if (status) {
      return (<strong className="me-auto" style={{color:'green'}}>Read</strong>)
    }else{
      return (<strong className="me-auto" style={{color:'red'}}>Unread</strong>)
    }
  }


  const renderNoteType = (notification)=>{
    if (notification.comment_author_name!==""){
      return <Toast.Body><span style={{color:'red'}}>{notification.comment_author_name}</span> commented your post <span style={{color:'red'}}>{notification.title}</span></Toast.Body>
    }else{
      return <Toast.Body><span style={{color:'red'}}>{notification.post_author_name}</span> made a post <span style={{color:'red'}}>{notification.title}</span> on your course<span style={{color:'red'}}> {notification.course_number}</span></Toast.Body>
    }
  }

  const handleClickStatus = (status, id) => {
    if (status === 0) {
      updateStatus(id);
    }
  }

  const updateStatus = async(id)=>{
    const res = await fetch('http://127.0.0.1:5005/api/v1/notification/'+id, {
      method: 'PUT',
      headers: {
        'accept': 'application/json',
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body:JSON.stringify({status: 1})
    });
    const data = await res.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      console.log(data.data);
      getNotifications();
    }
  }
  return (
    <>
      <img src={logo} alt="notifications" onClick={getNotifications} style={{width:'35px', pointerEvents:'stroke'}}/>
      <Offcanvas show={show} onHide={handleClose} {...props} placement='end'>
        <Offcanvas.Header closeButton>
          <Offcanvas.Title>Notifications</Offcanvas.Title>
        </Offcanvas.Header>
        <Offcanvas.Body>
        {notifications !== '' && notifications !== null && notifications !== undefined && 
        notifications.map((notification, index)=>(
          <Toast key={index} onClick={()=>handleClickStatus(notification.status, notification.id)}>
          <Toast.Header closeButton={false}>
            {renderStatus(notification.status)}
            <small>{notification.comment_time}</small>
          </Toast.Header>
          {renderNoteType(notification)}
        </Toast>
        ))}
        </Offcanvas.Body>
      </Offcanvas>
    </>
  );
}
export default OffCanvasExample;