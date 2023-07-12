import Form from 'react-bootstrap/Form'
import './index.css'
import {Link} from 'react-router-dom'
import React from 'react'

function Register(props) {
  const [email, setEmail] = React.useState('');
  const [pwd, setPwd] = React.useState('');
  const [cpwd, setCpwd] = React.useState('');
  const [name, setName] = React.useState('');
  const [authority, setAuthority] = React.useState(0);

  const handleAuthorityChange = (event) => {
    // console.log(event.target.value);
    setAuthority(parseInt(event.target.value));
  };

  const registerBtn = async () => {
    // console.log(authority);
    if (pwd === cpwd) {
      const response = await fetch('http://localhost:5005/api/v1/user/register', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
      },
      body: JSON.stringify({
        email: email,
        password: pwd,
        nickname: name,
        authority: authority,
        })
      });
      const data = await response.json();
      if (data.status !== 200) {

        alert(data.msg);
      } else {
        props.setTokenFn(data.token);
        localStorage.setItem('token', data.data.token);
        localStorage.setItem('owner', email);
        localStorage.setItem('authority', data.data.user.authority);
        localStorage.setItem('name', data.data.user.nickname);
        localStorage.setItem('avatar', data.data.user.Avatar);
        alert('Succeed!');
      }
    } else {
      alert('The two passwords you typed do not match.')
    }
  };
  return (
    <>
    <div className="account">
      <h2>Welcome back</h2>
      <form method="post">
        <div className="form-group">
          <label>Username</label><br/>
          <input type="text" name="username" placeholder="username" className='form-control' onChange={(event) => setEmail(event.target.value)} value={email}/>
        </div>
        <div className="form-group">
          <label>Password</label><br/>
          <input type="password" name="password" placeholder="password" className='form-control' onChange={(event) => setPwd(event.target.value)} value={pwd}/>
        </div>
        <div className="form-group">
          <label>Confirm password</label><br/>
          <input type="password" name="confirm-password" placeholder="confirm password" className='form-control' onChange={(event) => setCpwd(event.target.value)} value={cpwd}/>
        </div>
        <div className="form-group">
          <label>Nickname</label><br/>
          <input type="text" name="nickname" placeholder="nickname" className='form-control' onChange={(event) => setName(event.target.value)} value={name}/>
        </div>
        <div className="form-group">
          <label>Job</label><br/>
          <Form.Select className='job' id='authority' onChange={handleAuthorityChange} value={authority}>
            <option value="0">Student</option>
            <option value="1">Teacher</option>
          </Form.Select><br/>
        </div>
      </form>
      <div className="button_wrapper">
        <button type="submit" className="btn btn-primary sign"  onClick={registerBtn}>Sign up</button>
      </div>
      <div className='linkGroup'>
        <span>You already have an account?</span>
        <Link to="/login" className="link signLink">Sign in</Link>
      </div>
    </div>
    </>
  )
}

export default Register