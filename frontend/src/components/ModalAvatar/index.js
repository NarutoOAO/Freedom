import { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import './style.css';

function ModalAvatar(props) {
  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  function uploadFile() {
    const fileInput = document.getElementById('fileInput');
    const file = fileInput.files[0];
    const token = sessionStorage.getItem('token');

    const formData = new FormData();
    formData.append('file', file);

    fetch('http://127.0.0.1:5005/api/v1/user/avatar', {
      method: 'POST',
      headers: {
        'Authorization': token,
      },
      body: formData,
    })
      .then(response => response.json())
      .then((data) => {
        if (data.status !== 200) {
          alert(data.msg);
        } else {
          alert('Succeed!');
          // console.log(data.data);
          props.setAvatarFn(data.data.Avatar);
          sessionStorage.setItem('avatar', data.data.Avatar);
          window.location.reload();
          // console.log(sessionStorage.getItem('avatar'));
          handleClose();
        }
      }) 
      .catch((error) => {
        alert("Failed!")
      });
  };

  return (
    <>
      <Button variant="primary" onClick={handleShow} className='avatarModalBtn'>
        Change avatar
      </Button>

      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Avatars</Modal.Title>
        </Modal.Header>
        <Modal.Body className='modalBody'>
        <form encType="multipart/form-data">
          <input type="file" id="fileInput" />

        </form>
        </Modal.Body>
         <Modal.Footer>
           <Button variant="secondary" onClick={handleClose}>
             Close
           </Button>
           <Button variant="primary" type="button" onClick={uploadFile}>Upload</Button>
        </Modal.Footer>
      </Modal>
    </>
  );
}

export default ModalAvatar;
