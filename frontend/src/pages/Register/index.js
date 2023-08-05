import Form from 'react-bootstrap/Form'
import './index.css'
import {Link} from 'react-router-dom'
import React from 'react'
// define the register page
function Register(props) {
  const [email, setEmail] = React.useState('');
  const [pwd, setPwd] = React.useState('');
  const [cpwd, setCpwd] = React.useState('');
  const [name, setName] = React.useState('');
  const [authority, setAuthority] = React.useState(0);
  const [studyOption, setStudyOption] = React.useState('NULL');
  // regular expression for email and password
  const emailReg = /^([A-Za-z0-9_\-.])+@([A-Za-z0-9_\-.])+\.([A-Za-z]{2,4})$/;
  const passwordReg = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{6,16}$/;
  // handle the change of email, password, confirm password, name, authority and study option
  const handleAuthorityChange = (event) => {
    // console.log(event.target.value);
    setAuthority(parseInt(event.target.value));
    setStudyOption('NULL');
  };
  const handleStudyOptionChange = (event) => {
    setStudyOption(event.target.value);
  };
  
//define the register function
  const registerBtn = async () => {
    if (authority===0 && studyOption==="NULL"){
      alert('Student need choose its field of study')
      return
    }    if(!emailReg.test(email)){
        alert("The email is invalid!");
    }else if(!passwordReg.test(pwd)){
        alert("The password should include 6-16 bytes with at least 1 uppercase and 1 lowercase and 1 number!");
    }else if (pwd === cpwd) {
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
        studyoption:studyOption,
        })
      });
      const data = await response.json();
      if (data.status !== 200) {

        alert(data.msg);
      } else {
        props.setTokenFn(data.token);
        sessionStorage.setItem('token', data.data.token);
        sessionStorage.setItem('owner', email);
        sessionStorage.setItem('authority', data.data.user.authority);
        sessionStorage.setItem('name', data.data.user.nickname);
        sessionStorage.setItem('avatar', data.data.user.Avatar);
        sessionStorage.setItem('studyOption', data.data.user.studyoption);
        //console.log(data)
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
          </Form.Select>
        </div>
        {authority === 0 && (
            <div className="form-group">
              <label>Study Option</label>
              <Form.Select className="study-option" id="studyOption" onChange={handleStudyOptionChange} value={studyOption}>
                <option value="NULL">Field of study</option>
                <option value="AI">AI</option>
                <option value="IT">IT</option>
              </Form.Select>
            </div>
          )}<br/>
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