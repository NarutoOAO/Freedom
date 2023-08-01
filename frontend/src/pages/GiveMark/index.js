import './style.css';
import GradeReport from '../../components/GradeReport';
import { Button } from "antd";
import Accordion from 'react-bootstrap/Accordion';
import { useEffect, useState } from 'react';
export default function GiveMark(props) {
  const token = sessionStorage.getItem('token');
  const courseNumber = props.courseNumber;
  const [groups, setGroups] = useState(null);
  const authority = sessionStorage.getItem('authority');
  
  useEffect(()=>{
    getGroups();
  },[props.option])

  const getGroups = async()=>{
    const tempGroups = [];
    const res = await fetch('http://localhost:5005/api/v1/assignment_group/'+courseNumber,{
      method: "GET",
      headers: {
        'Authorization': token
      }
    })
    const data = await res.json();
    if(data.status!==200){
      alert(data.msg)
    }else{
      if(data.data!==null){
        for(let i=0;i<data.data.length;i++){
          let group_temp = data.data[i];
          // console.log(group_temp);
          try {
            const response = await fetch('http://127.0.0.1:5005/api/v1/assignment_mark/' + group_temp.id, {
              method: 'GET',
              headers: {
                'Content-Type': 'application/json',
                'Authorization': token,
              },
            });
      
            if (response.status === 200) {
              const data = await response.json();
              if(data.status===200){
                // console.log(data);
                group_temp['son'] = data.data;
              }else{
                alert('Failed!')
              }
            } else {
              throw new Error('Failed to fetch assignment');
            }
          } catch (error) {
            console.error(error);
          }
          tempGroups.push(group_temp)
        }
      }
    }
    setGroups(tempGroups)
  }

  const handleMarkChange=(event, index,index_new)=>{
    let groups_temp = groups;
    groups_temp[index].son[index_new].score = event.target.value;
    setGroups(groups_temp);
  }

  const handleContentChange=(event, index,index_new)=>{
    let groups_temp = groups;
    groups_temp[index].son[index_new].content = event.target.value;
    setGroups(groups_temp);
  }

  const submitMark=async(index,index_new)=>{
    const markValue = parseInt(groups[index].son[index_new].score);
    const max_score = groups[index].son[index_new].max_score;
    const content = groups[index].son[index_new].content
    // Check if markValue is a valid number and within the valid range
    if (isNaN(markValue) || markValue <= 0 || markValue > max_score) {
      alert('Please enter a valid mark');
      return;
    }
    const markInfotmaion={
      ass_mark_id: groups[index].son[index_new].ass_mark_id,
      mark: markValue,
      content: content,
    }
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/assignment_grade', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify(markInfotmaion),
      });
  
      if (response.status === 200) {
        alert('Scores have been marked successfully.');
      } else {
        throw new Error('Failed to Publish');
      }
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <div>
    {parseInt(authority)===0 && <div className="enroll-course-container" style={{marginLeft:'0', marginRight:'90px'}}>
        <h2>Check yout grade</h2>
      <GradeReport/>
      </div>}
    <div>
      <h2>Group Mark</h2>
      { groups!==null && groups!== '' &&
        groups.map((group,index)=>(
          <Accordion defaultActiveKey="0" style={{marginTop:"20px", width:'95%'}} key={index}>
          <Accordion.Item eventKey={index}>
            <Accordion.Header>{group.group_name}</Accordion.Header>
            <Accordion.Body>
            <table className="table table-hover">
            <colgroup>
              <col style={{ width: '10%' }} />
              <col style={{ width: '30%' }} />
              <col style={{ width: '15%' }} />
              <col style={{ width: '20%' }} />
              <col style={{ width: '15%' }} />
              <col style={{ width: '15%' }} />
            </colgroup>
            <thead>
              <tr>
                <th scope="col">Student Name</th>
                <th scope="col">Link</th>
                <th scope="col">Mark</th>
                <th scope="col">Feedback</th>
                <th scope="col">Group</th>
                <th scope="col">Button</th>
              </tr>
            </thead>
            {group.son !== null && group.son!=='' &&
              group.son.map((son,index_new) => (
            <tbody>
                <tr key={index_new}>
                  <td>{son.file_url.substring(son.file_url.lastIndexOf("/") + 1, son.file_url.lastIndexOf("."))}</td>
                  <td>
                    <a href={son.file_url} target="_blank" rel="noopener noreferrer">
                      {son.file_url}
                    </a>
                  </td>
                  <td>
                    <input type="text" placeholder={son.score!==0?son.score:"mark"} onChange={(event) => handleMarkChange(event, index,index_new)}/>
                  </td>
                  <td>
                    <input type="text" placeholder={son.content!==''?son.content:"feedback"} onChange={(event) => handleContentChange(event, index,index_new)}/>
                  </td>
                  <td>{son.group_name}</td>
                  <td>
                    <Button style={{ backgroundColor: "#87CEEB" }}  onClick={() =>submitMark(index,index_new)}>Mark</Button>
                  </td>
                </tr>
            </tbody>))}
          </table>
            </Accordion.Body>
          </Accordion.Item>
        </Accordion>))}
      </div>
    </div>
  )
}
