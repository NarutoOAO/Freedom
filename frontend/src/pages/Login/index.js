import '../Register/index.css'
import {Link} from 'react-router-dom'
import React from 'react'
function Login(props) {
  const [email, setEmail] = React.useState('');
  const [pwd, setPwd] = React.useState('');

  const loginBtn = async () => {
    const response = await fetch('http://localhost:5005/api/v1/user/login', {
      method: 'POST',
      headers: {
        'accept': 'application/json',
        'Content-type': 'application/json',
      },
      body: JSON.stringify({
        email: email,
        password: pwd,
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      // console.log(data);
      props.setTokenFn(data.data.token);
      localStorage.setItem('token', data.data.token);
      localStorage.setItem('owner', email);
      localStorage.setItem('authority', data.data.user.authority);
      localStorage.setItem('name', data.data.user.nickname);
      localStorage.setItem('avatar', data.data.user.Avatar);
      alert("Succeed!");
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
        </div><br/>
      </form>
      <div className="button_wrapper">
        <button type="submit" className="btn btn-primary sign"  onClick={loginBtn}>Sign in</button>
      </div>
      <div className='linkGroup'>
        <span>You don't have an account?</span>
        <Link to="/register" className="link signLink">Sign up</Link>
      </div>
    </div>
    </>
  )
}

export default Login