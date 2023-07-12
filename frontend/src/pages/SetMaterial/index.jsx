import React, {useEffect, useState,useRef} from 'react';
import ModalMaterial from '../../components/ModalMaterial';
import {Button} from 'antd';
import "./setMaterial.css";
import PostMaterial from '../../components/Material';

import { useParams } from 'react-router-dom';

export default function SetMaterial () {
  const authority=localStorage.getItem('authority');
  const inf = useRef();
  const { courseNumber } = useParams();
  const[weekInfo,setWeekInfo]=useState([
    {
      id:1,
      name:"week 1"
    },
    {
      id:2,
      name:"week 2"
    },
    {
      id:3,
      name:"week 3"
    },
    {
      id:4,
      name:"week 4"
    }
    ,
    {
      id:5,
      name:"week 5"
    }
    ,
    {
      id:6,
      name:"week 6"
    }
    ,
    {
      id:7,
      name:"week 7"
    }
    ,
    {
      id:8,
      name:"week 8"
    }
    ,
    {
      id:9,
      name:"week 9"
    }
    ,
    {
      id:10,
      name:"week 10"
    }
  ]);

  function handleKeyUp(e){
    //if do search function, delete below
    if(e.keyCode!==13)return
    const weekinfoObj={id:Date.now(),name:e.target.value};
    setWeekInfo([...weekInfo,weekinfoObj]);
  }

  return (
    <div className = "SetMaterial-container">
      <div className="input-Mcontainer">
        <input type="text" name="inf" ref={inf} placeholder="Enter your search information" onKeyUp={handleKeyUp}/>
        <Button >Search</Button>
      </div>
      {authority !== '0' && (
      <div className="postMaterial-container">
            <PostMaterial courseNumber={courseNumber} />
      </div>
      )}
      <div className="weekInfo-container">
      {
        weekInfo.map((item)=>{
         return(<ModalMaterial key={item.id} modelName={item.name} courseNumber={courseNumber}/>)
         
          
        })
      }
      </div>
    </div>
  )
}