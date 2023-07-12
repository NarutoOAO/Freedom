import React, { useState, useEffect } from 'react';
import { Collapse } from "antd";
import "./index.css";
import logo0 from '../../images/pdf-icon.svg'
import logo1 from '../../images/ppt-icon.svg'

function ModalMaterial(props) {
  const { Panel } = Collapse;
  const courseNumber = props.courseNumber;
  const weekNumber = props.modelName;
  const token = localStorage.getItem('token');
  const [postMaterial, setPostMaterial] = useState([]);
  const authority=localStorage.getItem('authority');

  const fetchMaterial = async () => {
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/material/' + courseNumber + '/' + weekNumber, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
      });

      if (response.status === 200) {
        const data = await response.json();
        setPostMaterial(data.data);
      } else {
        throw new Error('Failed to fetch material');
      }
    } catch (error) {
      console.error(error);
    }
  };


  const publish_function = async (props) => {
    const requestPublish = {
      material_id: props ,
      publish: parseInt('1'),
    };
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/material' , {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify(requestPublish),
      });

      if (response.status === 200) {
        fetchMaterial();
      } else {
        throw new Error('Failed to Publish');
      }
    } catch (error) {
      console.error(error);
    }
  };
  const unpublish_function = async (props) => {
    const requestUnpublish = {
      material_id: props ,
      publish: parseInt('2'),
    };
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/material' , {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify(requestUnpublish),
      });

      if (response.status === 200) {
        fetchMaterial();
      
      } else {
        throw new Error('Failed to Publish');
      }
    } catch (error) {
      console.error(error);
    }
  };
  const delete_function = async (props) => {
    const requestDelete= {
      material_id: props ,
    };
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/material' , {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify(requestDelete),
      });

      if (response.status === 200) {
        fetchMaterial();
      
      } else {
        throw new Error('Failed to Publish');
      }
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    fetchMaterial();
  }, []);
  //ues material.status to change
  const getStatusText = (type) => {
    return type === 2 ? 'Unpublish' : 'Publish';
  };
  const getButtonColor = (type) => {
    return type === 2 ? '#fbc02d' : '#7ed957';
  };
  return (
    <div className='material'>
      <Collapse defaultActiveKey={['1']} className='panel'>
        <Panel header={props.modelName} key="0" className="weekModule" onClick={fetchMaterial}>
          <div classname='panel-content' className='element'>   
          {postMaterial && postMaterial.length > 0 ? (
            postMaterial.map((material) => (
              <div key={material.file_name}>
                {authority !== '0' && (
                <span>
                  <button className="status-button"  style={{ backgroundColor: getButtonColor(material.publish) }}>{getStatusText(material.publish)}</button>
                </span>
                )}
                {material.type === 0 ? (
                    <span><img src={logo0} alt="PDF" /></span>
                  ) : (
                    <span><img src={logo1} alt="PPT" /></span>
                  )}
                  
                <span> {material.file_name} </span>
                <span><a href={material.file_url} target="_blank" rel="noopener noreferrer">{material.file_url}</a> </span>
                {authority !== '0' && (
                <span className="function-content">
                  <button className="function-button" style={{ backgroundColor: '#cfd8dc' }} onClick={() => publish_function(material.material_id)}>Publish</button>
                  <button className="function-button" style={{ backgroundColor: '#eceff1' }} onClick={() => unpublish_function(material.material_id)}>Unpublish</button>
                  <button className="function-button" style={{ backgroundColor: '#ef9a9a' }} onClick={() => delete_function(material.material_id)}>Delete</button>
                </span>
                )}
                <hr className="line-separator" />
              </div>
            ))
          ) : (
            <p>No materials found.</p>
          )}
          </div>      
        </Panel>
      </Collapse>
    </div>
  );
}

export default ModalMaterial;
