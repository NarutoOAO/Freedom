import { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
// define a modal to reset the password
function ModalReset() {
  const [show, setShow] = useState(false);
  const [currentPassword, setCurrentPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [confirmNewPassword, setConfirmNewPassword] = useState('');
// close the modal
  const handleClose = () => {
    setShow(false);
  };
// open the modal
  const handleShow = () => setShow(true);
// define the function to reset the password
  const handleResetPassword = async () => {
    const apiUrl = 'http://127.0.0.1:5005/api/v1/user/password';
    const token = sessionStorage.getItem('token');
    const requestBody = {
      password: currentPassword,
      new_password: newPassword,
      confirm_new_password: confirmNewPassword,
    };

    try {
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token
        },
        body: JSON.stringify(requestBody),
      });

      const data = await response.json();

      if (data.status === 200) {
        // Password reset successful
        console.log(data);
        alert('Password reset successful');
      } else {
        // Password reset failed
        alert(data.msg);
      }
    } catch (error) {
      console.log(error);
      alert('An error occurred while resetting the password.');
    }

    // Reset the form fields
    setCurrentPassword('');
    setNewPassword('');
    setConfirmNewPassword('');

    // Close the modal
    handleClose();
  };

  return (
    <>
      <Button variant="primary" onClick={handleShow}>
        Reset password
      </Button>

      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Reset password</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group className="mb-3" controlId="currentPassword">
              <Form.Label>Current password</Form.Label>
              <Form.Control
                type="password"
                value={currentPassword}
                onChange={(e) => setCurrentPassword(e.target.value)}
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="newPassword">
              <Form.Label>New password</Form.Label>
              <Form.Control
                type="password"
                value={newPassword}
                onChange={(e) => setNewPassword(e.target.value)}
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="confirmNewPassword">
              <Form.Label>New password confirm</Form.Label>
              <Form.Control
                type="password"
                value={confirmNewPassword}
                onChange={(e) => setConfirmNewPassword(e.target.value)}
              />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" onClick={handleResetPassword}>
            Submit
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
}

export default ModalReset;
