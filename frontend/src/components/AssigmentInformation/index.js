import React, { useState, useEffect } from 'react';
import "./index.css";
import logo0 from '../../images/pdf-icon.svg'
import SetAssigementEndTime from "../SetAssigementEndTime"
import StudentUploadAssigment from "../StudentUploadAssigment"
import MarkAssigment from "../MarkAssigment"
import StudentDeleteSumbitAssigment from "../StudentDeleteSumbitAssigment"

function AssigmentInformation(props) {
  //console.log(props)
  const courseNumberUseForAssigment = props.courseNumber;
  const token = sessionStorage.getItem('token');
  const [postAssigment, setPostAssigment] = useState([]);
  const authority=sessionStorage.getItem('authority');
  
  const fetchAssigment = async () => {

    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/assignment/' + courseNumberUseForAssigment, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
      });

      if (response.status === 200) {
        const data = await response.json();
        setPostAssigment(data.data);
      } else {
        throw new Error('Failed to fetch assigment');
      }
    } catch (error) {
      console.error(error);
    }
  };

  

  
  const delete_assigment = async (props) => {
    const requestDelete= {
      assignment_id: props ,
    };
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/assignment' , {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify(requestDelete),
      });

      if (response.status === 200) {
        fetchAssigment();
        //props.setAssigmentFlag(1);
      } else {
        throw new Error('Failed to Publish');
      }
    } catch (error) {
      console.error(error);
    }
  };

  const delete_sumbit_assigment = async (event) => {
      //console.log('1')
      const requestAssignmentDelete=
      {
        ass_mark_id: parseInt(event),
      };
      console.log(requestAssignmentDelete)
      
       
        try {
          const response = await fetch('http://127.0.0.1:5005/api/v1/assignment_solution' , {
            method: 'DELETE',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': token,
            },
            body: JSON.stringify(requestAssignmentDelete),
          });

          if (response.status === 200) {
            alert('Success Delete')
          } else {
            throw new Error('Failed to Delete');
          }
        } catch (error) {
          console.error(error);
        }
     
  };
  
 
  

 
  useEffect(() => {
    fetchAssigment();
    props.setAssigmentFlag(0);
    // eslint-disable-next-line
  },[props.assigmentFlag,props.option]);
  
  const getStatusText = (time) => {
    return time === '0001-01-01 00:00:00' ? 'Unpublish' : 'Publish';
  };
  const getButtonColor = (time) => {
    return time === '0001-01-01 00:00:00' ? '#fbc02d' : '#7ed957';
  };
  return (
    <div className='assigment'>
      <div className='assigment-panel-content' style={{ maxHeight: 'calc(100vh - 200px)', overflowY: 'auto' }}>
        {postAssigment && postAssigment.length > 0 ? (
          postAssigment.map((assigment,index) => (
            <div key={index}>
               {authority === '1' ? (
                <div>
                  <div style={{ backgroundColor: '#f5f5f5', padding: '10px', borderRadius: '4px', fontWeight: 'bold' }}>
                    <span>Assignment Name: </span>
                    <span>{assigment.file_name}</span>
                    <span style={{marginLeft:'20px'}}>Mark: {assigment.max_score}</span>
                  </div>
                  
                    <span>
                      <button className="assigment-status-button" style={{ backgroundColor: getButtonColor(assigment.end_time) }}>
                        {getStatusText(assigment.end_time)}
                      </button>
                    </span>
              
                  <span><img src={logo0} alt="PDF" /></span>
                  <span><a href={assigment.file_url} target="_blank" rel="noopener noreferrer">{assigment.file_url}</a></span>
                  
                    <span className="assigment-function-content">
                      <SetAssigementEndTime assignment_id={assigment.assignment_id} props={props}/>
                      <button className="assigment-function-button" style={{ backgroundColor: '#ef9a9a' }} onClick={() => delete_assigment(assigment.assignment_id)}>Delete</button>
                      <MarkAssigment assignment_id={assigment.assignment_id} />
                    </span>
                  <hr className="assigment-line-separator" />
                </div>
               
               ) : authority === '0' && assigment.end_time !== '0001-01-01 00:00:00' && (
                <div>
                  <div style={{ backgroundColor: '#f5f5f5', padding: '10px', borderRadius: '4px', fontWeight: 'bold' }}>
                    <span>Assignment Name: </span>
                    <span>{assigment.file_name}</span>
                  </div>
                  <span><img src={logo0} alt="PDF" /></span>
                  <span><a href={assigment.file_url} target="_blank" rel="noopener noreferrer">{assigment.file_url}</a></span>
                  <span className="assigment-function-content">
                    <StudentUploadAssigment assignment_id={assigment.assignment_id} end_time={assigment.end_time}/>
                    <StudentDeleteSumbitAssigment assignment_id={assigment.assignment_id} />
                  </span>
                </div>
              ) }
            </div>
          ))
        ) : (
          <p>No assignment found.</p>
        )}
      </div>
     
    </div>

  
  );

}

export default AssigmentInformation;
