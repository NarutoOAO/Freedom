import React, { useEffect } from 'react'
import Button from 'react-bootstrap/Button';
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import {Link} from 'react-router-dom'
import './style.css'
function NavScrollExample(props) {
  
  const authority = parseInt(localStorage.getItem('authority'));
  function logout() {
    localStorage.removeItem('token');
    localStorage.removeItem('owner');
    localStorage.removeItem('authority');
    localStorage.removeItem('avatar');
    localStorage.removeItem('name');
    props.setTokenFn(null);
    document.body.style.background='linear-gradient(800deg,#e3c5eb,#a9c1ed)';
  }
  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);
  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <Navbar className='nav' id='navView' >
      <Container fluid>
        <Navbar.Brand href="/" className='left'>FREEDOM</Navbar.Brand>
        <Navbar.Toggle aria-controls="navbarScroll" />
        <Navbar.Collapse id="navbarScroll">
        {localStorage.getItem('token') && (<Nav
            className="me-auto my-2 my-lg-0 navBar"
            style={{ maxHeight: '100px' }}
            navbarScroll
          >
          <Link to={authority?'/teacher/dashboard':'/student/dashboard'} className='tab-link dashboard-link'><MenuItem onClick={handleClose}>Dashboard</MenuItem></Link>
          <Link to={authority?'/create_course':'/enroll_course'} className='tab-link dashboard-link'><MenuItem onClick={handleClose}>Courses</MenuItem></Link>
            {/* <Form className="d-flex">
            <Form.Control
              type="search"
              placeholder="Search"
              className="me-2 search"
              aria-label="Search"
            />
            <Button variant="outline-success">Search</Button>
          </Form> */}
          <div>
                <Button
                  id="basic-button"
                  className='dropDown'
                >
                  <img className='AccountCircleIcon' src={localStorage.getItem('avatar')} alt='avatar'
                  onClick={handleClick}/>
                  <span className='nickName'>{localStorage.getItem('name')}</span>
                </Button>
                <Menu
                  id="basic-menu"
                  anchorEl={anchorEl}
                  open={open}
                  onClose={handleClose}
                  MenuListProps={{
                    'aria-labelledby': 'basic-button',
                  }}
                >
                  <Link to='/profile' className='tab-link'><MenuItem onClick={handleClose}>Profile</MenuItem></Link>
                  <hr/>
                  <Link to='/' id="logout" className='tab-link'><MenuItem onClick={logout}>Logout</MenuItem></Link>
                </Menu>
              </div>
          </Nav>
          )}
        </Navbar.Collapse>
        
      </Container>
    </Navbar>
  );
}

export default NavScrollExample;