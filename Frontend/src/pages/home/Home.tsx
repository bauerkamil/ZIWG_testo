import React from 'react';
import { Outlet } from 'react-router-dom';
import './Home.module.scss';
import axios from 'axios';
import {Button} from "@/components/ui/button.tsx";

interface ResponseType {
    Token: string;
}

const Home: React.FC = () => {
    const fetchToken = async () => {
        const response = await axios.get<ResponseType>("https://func-pwr-testo-dev.azurewebsites.net/api/auth");
        console.log(response.data.Token);
    }

  return (
    <div>
      <h1>Welcome to the Homepage</h1>
      <p>This is the content of the homepage.</p>
        <Button onClick={() => fetchToken()}>Kliknij mnie</Button>
      <Outlet />
    </div>
  );
};

export default Home;