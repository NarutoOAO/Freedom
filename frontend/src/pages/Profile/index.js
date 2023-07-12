import React, { useEffect, useState } from 'react';
import './style.css'
import ModalReset from '../../components/ModalReset'
import ModalName from '../../components/ModalName'
import ModalAvatar from '../../components/ModalAvatar';
export default function Profile() {
  const [name, setName] = useState(localStorage.getItem('name'));
  const email = localStorage.getItem('owner');
  const [avatar, setAvatar] = useState(localStorage.getItem('avatar'));

  return (
<div className="profile-container">
  <div className='avatar'>
    <img src={avatar} alt="avatar" className="profileAvatar" />
  </div>

  <div>
    <span>Nickname:</span> {name}
  </div>
  <div>
    <span>Email:</span> {email}
  </div>
  <div className='profileBtnGroup'>
    <ModalAvatar setAvatarFn={setAvatar}/>
    <ModalReset/>
    <ModalName setNameFn={setName}/>
  </div>
</div>

  )
}
