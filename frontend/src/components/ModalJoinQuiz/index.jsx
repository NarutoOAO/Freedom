import React, { useEffect, useState } from 'react';
// import { useAlert } from '../../hooks/useAlert';
import { Card, Typography, List } from 'antd';
import { TextField } from '@mui/material';
import { Button, Modal } from 'react-bootstrap';


const { Title } = Typography;

export default function ModalJoinQuiz(props) {
    const token = sessionStorage.getItem('token');
    const [showModal, setShowModal] = useState(false);
    const [singleQuiz, setSingleQuiz] = useState([]);
    // const { alert, Alert } = useAlert();
    const quiz_id = props.quiz_id;
    const [answer, setAnswer] = useState({});

    const handleNavigation = async () => {
        
        setShowModal(true);
        try {
            const response = await fetch('http://127.0.0.1:5005/api/v1/quiz_question/' + quiz_id, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': token,
                },
            });

            if (response.status === 200) {
                const data = await response.json();
                if (data.status === 200) {
                    console.log(data.data)
                    setSingleQuiz(data.data);
                    
                } else {
                    alert(data.msg);
                }
            } else {
                throw new Error('Failed to fetch posts');
            }
        } catch (error) {
            alert(error);
        };
    };
    useEffect(() => {
        if (props.showModal) {
            handleNavigation();
        }
    }, [props.showModal]);

    const handleCloseModal = () => {
        props.onClose();
        setAnswer({});
        setShowModal(false);
    };
    const handleAnswerChange = (event, quizAnswerId) => {
        setAnswer(prevState => ({
          ...prevState,
          [quizAnswerId]: event.target.value
        }));
      };

    const submitFunc = (quiz_question_id,answer,question_number) => {
        
        const request = {
            quiz_question_id: parseInt(quiz_question_id),
            user_answer: answer[question_number],
          };
          fetch('http://127.0.0.1:5005/api/v1/quiz_mark', {
            method: 'POST',
            headers: {
              'Authorization': token,
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(request),
          })
            .then(response => response.json())
            .then((data) => {
              if (data.status !== 200) {
                //console.log('1')
                alert("Failed!")
              } else {
                alert('Answer Post');
                
                setAnswer(prevState => ({
                  ...prevState,
                  [question_number]: ''
                }));

              }
            })
           ;
        //handleCloseModal();
    };
    return (
        <>
            {/* Display the alert */}
            {/* {Alert} */}

            <Modal show={showModal} onHide={handleCloseModal}>
                <Modal.Header closeButton>
                    <Modal.Title>Complete Quiz</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                <div style={{ color: 'gray', fontStyle: 'italic' ,marginTop:'10px'}}>
                    <p>For single Question, write 'A'</p>
                    <p>For multiple Question, if the answer is A and B, write 'AB'</p>
                </div>
                {singleQuiz && singleQuiz.map((singlequiz, index) => {
                        const quizAnswerId = singlequiz.question_number;

                        return (
                        <Card className="quiz-card" key={index} style={{border:"2px solid",marginBottom:"10px"}}>
                            <Title level={4}>Question {singlequiz.question_number}{singlequiz.type === 1 && " (Single choice question)"}{singlequiz.type === 0 && " (Multiple choice question)"} </Title>
                            <h6>Question Mark: {singlequiz.score}</h6>
                            <div className="description"><h6>Question: {singlequiz.description}</h6></div>
                            <List>
                            <List.Item>A: {singlequiz.select_A}</List.Item>
                            <List.Item>B: {singlequiz.select_B}</List.Item>
                            <List.Item>C: {singlequiz.select_C}</List.Item>
                            <List.Item>D: {singlequiz.select_D}</List.Item>
                            </List>
                            <TextField
                            label="Answer"
                            margin="normal"
                            variant="outlined"
                            value={answer[quizAnswerId] || ''}
                            onChange={(event) => handleAnswerChange(event, quizAnswerId)}
                            />
                            <br /><br />
                            <Button onClick={() => submitFunc(singlequiz.quiz_question_id, answer,singlequiz.question_number)}>Submit</Button>
                        </Card>
                        );
                    })}
                    </Modal.Body>
                <Modal.Footer>
                    {/* <Button onClick={submitFunc}>Submit</Button> */}
                </Modal.Footer>
            </Modal>
        </>
    );
}