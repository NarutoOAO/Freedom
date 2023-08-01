import './index.css'
import Button from 'react-bootstrap/Button';
import React from 'react';
import {Link} from 'react-router-dom'
function Layout() {
  
  return (
    <div className='layoutBody'>
      <h1>Welcome to Freedom!</h1>
      <div className='container'>
      <div className='button-wrapper'>
        <Button variant="outline-primary"><Link to="/login" className="link">Sign in</Link></Button>
        <Button variant="outline-danger"><Link to="/register" className="link">Sign up</Link></Button>
      </div>
    </div>
    </div>
  )
}

export default Layout