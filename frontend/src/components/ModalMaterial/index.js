import React, { useState, useEffect } from 'react';
import { Collapse } from "antd";
import "./index.css";
import logo0 from '../../images/pdf-icon.svg'
import logo1 from '../../images/ppt-icon.svg'
// define a component to show the material
function ModalMaterial(props) {
  // Destructure the Collapse component from the 'antd' library
  const { Panel } = Collapse;
  const courseNumber = props.courseNumber;
  const weekNumber = props.modelName;
  // Get the token from sessionStorage
  const token = sessionStorage.getItem('token');
  // Use to store material data
  const [postMaterial, setPostMaterial] = useState([]);
  // Get the token from authority
  const authority=sessionStorage.getItem('authority');
  // Function to fetch materials for the selected week
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

  // Function to mark material as published
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

  // Function to mark material as unpublished
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

  // Function to delete material
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
    // eslint-disable-next-line
  }, []);
  // Function to get the status based on the publish type
  const getStatusText = (type) => {
    return type === 2 ? 'Unpublish' : 'Publish';
  };
  // Function to get the color based on the publish type
  const getButtonColor = (type) => {
    return type === 2 ? '#fbc02d' : '#7ed957';
  };
  return (
    <div className='material'>
      <Collapse defaultActiveKey={['1']} className='panel'>
        <Panel header={props.modelName} key="0" className="weekModule" onClick={fetchMaterial}>
          {/* Panel representing the week */}
          <div classname='panel-content' className='element'>   
          {postMaterial && postMaterial.length > 0 ? (
            postMaterial.map((material) => (
              <div key={material.file_name}>
                {/* show materials inforamtion if there are any */}
                {authority !== '0' && (
                <span>
                  <button className="status-button"  style={{ backgroundColor: getButtonColor(material.publish) }}>{getStatusText(material.publish)}</button>
                </span>
                )}
                {/* base on type to show different logo */}
                {material.type === 0 ? (
                    <span><img src={logo0} alt="PDF" /></span>
                  ) : (
                    <span><img src={logo1} alt="PPT" /></span>
                )}
                  
                <span> {material.file_name} </span>
                <span><a href={material.file_url} target="_blank" rel="noopener noreferrer">{material.file_url}</a> </span>
                 {/* only teacher can unpublish,publish nad delete material */}
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
            // if no material found, show below mseeage
            <p>No materials found.</p>
          )}
          </div>      
        </Panel>
      </Collapse>
    </div>
  );
}

export default ModalMaterial;
