import * as React from 'react';
import Accordion from '@mui/material/Accordion';
import AccordionSummary from '@mui/material/AccordionSummary';
import AccordionDetails from '@mui/material/AccordionDetails';
import Typography from '@mui/material/Typography';
import Infrastructure from './Infrastructure';
import generateRandomHexColor from '../../utils/randomColor';

const Home: React.FC = () => {
  return (
    <div>
      <Infrastructure color={generateRandomHexColor()} header={'S3'} dataObject={{ buckets: 3 }}>
        <Typography>S3 Buckets :- <div style={{ display: 'inline', color: 'red' }}>3</div> </Typography>
      </Infrastructure>
    </div>
  );
};

export default Home;
