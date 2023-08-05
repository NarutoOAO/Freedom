
import './style.css';
import { useParams } from 'react-router-dom';
import React, { useState, useEffect } from 'react';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
// define the grade report
export default function GradeReport() {
  const { courseNumber } = useParams();
  const [markAssigmentInfomtaion,setmMrkAssigmentInfomtaion]= useState([]);
  // Get the token from sessionStorage
  const token = sessionStorage.getItem('token');
  // State variable to control the visibility of the modal
  const [show, setShow] = useState(false);
  // State variable to trigger fetching of grade information
  const [flagForGrade,setFlagForGrade]=useState(false)
  const [markQuizInfomtaion,setQuizInfomtaion]= useState([]);
  
  // Function to handle the display of the modal
  const handleShow = async () => {
    fetchMarkAssigmentInfomtaion();
    sumQuizMark();
    setShow(true);
  }

  // Function to handle the closure of the modal
  const handleClose = () => {
    setShow(false);
  }

  useEffect(() => {
    if (flagForGrade) {
      fetchMarkAssigmentInfomtaion();
      sumQuizMark();
      setFlagForGrade(false); // Reset the flag to prevent further fetching
    }
     // eslint-disable-next-line
  }, [flagForGrade]);

  // Function to fetch assignment grade information from the API
  const fetchMarkAssigmentInfomtaion = async () => {
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/assignment_solution/' + parseInt(courseNumber), {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
      });

      if (response.status === 200) {
        const data = await response.json();
        setmMrkAssigmentInfomtaion(data.data);
      } else {
        throw new Error('Failed to fetch assigment');
      }
    } catch (error) {
      console.error(error);
    }
  };

  // Function to fetch quiz grade information from the API
  const sumQuizMark = async () => {
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/quiz_sum/' + parseInt(courseNumber), {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
      });

      if (response.status === 200) {
        const data = await response.json();
        setQuizInfomtaion(data.data);
      } else {
        throw new Error('Failed to fetch assigment');
      }
    } catch (error) {
      console.error(error);
    }
  };

  // Function to calculate letter grade
  const calculateLetterGrade = (percentage) => {
    if (percentage >= 85 && percentage <= 100) {
      return 'HD';
    } else if (percentage >= 75 && percentage <= 84) {
      return 'DN';
    } else if (percentage >= 65 && percentage <= 74) {
      return 'CR';
    } else if (percentage >= 50 && percentage <= 64) {
      return 'PS';
    } else {
      return 'FL';
    }
  };
  
  return (
    <>
    <button onClick={handleShow}>Check</button>
    <Modal show={show} onHide={handleClose} size='lg'>
      <Modal.Header closeButton>
        <Modal.Title> 
          <h2>Grade</h2>
        </Modal.Title>
     
      </Modal.Header>
      <Modal.Body className='modalBody' style={{  fontFamily: 'Arial, sans-serif', fontSize: '18px', color: '#333', backgroundColor: '#f9f9f9', padding: '10px', border: '1px solid #ccc', borderRadius: '5px', boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',maxHeight: '500px',overflowY: 'auto', }}>
        <div className="tableContainer">
          <table>
             {/* Table to display assignment and quiz grade information */}
            <thead>
              <tr>
                <th style={{ width: '20%' }}>Grade item</th>
                <th style={{ width: '5%' }}>Grade</th>
                <th style={{ width: '5%' }}>Range</th>
                <th style={{ width: '5%' }}>Percentage</th>
                <th style={{ width: '20%' }}>Letter grade</th>
                <th style={{ width: '65%' }}>Feedback</th>
              </tr>
            </thead>
            <tbody>
              {/* Mapping through quiz grade information to display rows */}
              {markQuizInfomtaion && markQuizInfomtaion.length > 0 && (
                markQuizInfomtaion.map((item, index) => {
                  // Calculate the percentage and letter grade for each quiz
                  const percentage = (item.Score / item.MaxScore) * 100;
                  const letterGrade = calculateLetterGrade(percentage);

                  return (
                    <tr key={index}>
                      <td>{item.QuizName}</td>
                      <td>{item.Score}</td>
                      <td>0 - {item.MaxScore}</td>
                      <td>{percentage.toFixed(2)}%</td>
                      <td>{letterGrade}</td>
                      <td>-</td>
                    </tr>
                  );
                })
              )}
              {/* Mapping through assignment grade information to display rows */}
              {markAssigmentInfomtaion && markAssigmentInfomtaion.length > 0 && (
                markAssigmentInfomtaion.map((item, index) => {
                  // Calculate the percentage and letter grade for each assignment
                  const percentage = (item.score / item.max_score) * 100;
                  const letterGrade = calculateLetterGrade(percentage);

                  return (
                    <tr key={index}>
                      <td>{item.ass_name}</td>
                      <td>{item.score}</td>
                      <td>0 - {item.max_score}</td>
                      <td>{percentage.toFixed(2)}%</td>
                      <td>{letterGrade}</td>
                      <td>{item.content}</td>
                    </tr>
                  );
                })
              )}
              {/* Calculate and display the total score and average percentage */}
              {(markQuizInfomtaion?.length || markAssigmentInfomtaion?.length) && (
                <tr>
                  <td>Total Score:</td>
                  <td>
                    {(
                      ((markQuizInfomtaion?.reduce((total, item) => total + (item.Score / item.MaxScore) * 100, 0) || 0) +
                      (markAssigmentInfomtaion?.reduce((total, item) => total + (item.score / item.max_score) * 100, 0) || 0))/
                      ((markQuizInfomtaion?.length || 0) + (markAssigmentInfomtaion?.length || 0))
                    ).toFixed(2)
                    }
                  </td>
                  <td>0-100</td>
                  <td>
                    {(
                      ((markQuizInfomtaion?.reduce((total, item) => total + (item.Score / item.MaxScore) * 100, 0) || 0) +
                      (markAssigmentInfomtaion?.reduce((total, item) => total + (item.score / item.max_score) * 100, 0) || 0))/
                      ((markQuizInfomtaion?.length || 0) + (markAssigmentInfomtaion?.length || 0))
                    ).toFixed(2)
                    }%
                  </td>
                  <td>{calculateLetterGrade(
                    ((markQuizInfomtaion?.reduce((total, item) => total + (item.Score / item.MaxScore) * 100, 0) || 0) +
                    (markAssigmentInfomtaion?.reduce((total, item) => total + (item.score / item.max_score) * 100, 0) || 0)) /
                    ((markQuizInfomtaion?.length || 0) + (markAssigmentInfomtaion?.length || 0))
                  )}</td>
                  <td>-</td>
                  <td>-</td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
    
      </Modal.Body>
          <Modal.Footer>
            <Button variant="primary" onClick={handleClose}>Close</Button>
          </Modal.Footer>
    </Modal>
    </>
  );
  
}
