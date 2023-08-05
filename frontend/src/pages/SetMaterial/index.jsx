import React, {useState} from 'react';
import ModalMaterial from '../../components/ModalMaterial';
import "./setMaterial.css";
import PostMaterial from '../../components/Material';
import SearchMaterial from '../../components/SearchMaterial';
import { useParams } from 'react-router-dom';

export default function SetMaterial () {
  // Get the authority from session storage 
  const authority=sessionStorage.getItem('authority');
  const { courseNumber } = useParams();
  // eslint-disable-next-line
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

  
  return (
    <div className = "SetMaterial-container">
      <div className="input-Mcontainer">
        <SearchMaterial courseNumber={courseNumber} />
      </div>
      {/* As teachers, they can post material*/}
      {authority !== '0' && (
      <div className="postMaterial-container">
            <PostMaterial courseNumber={courseNumber} />
      </div>
      )}
      {/*See material inforamtion*/}
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