import axios from "axios";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const TOKEN_LINK = import.meta.env.VITE_TOKEN_LINK;

interface IResponseTypeLogin {
  key: string | null;
  oauth_token: string | null;
  oauth_verifier: string | null;
}

interface IResponseTokenParams {
  Token?: string;
  FirstName?: string;
  LastName?: string;
}

const LoadingPage = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const urlParams = new URLSearchParams(window.location.search);

    if (!urlParams.get("key")) {
      return;
    }

    const responseObj: IResponseTypeLogin = {
      key: urlParams.get("key"),
      oauth_token: urlParams.get("oauth_token"),
      oauth_verifier: urlParams.get("oauth_verifier"),
    };

    if (responseObj) {
      const getData = async () => {
        try {
          const response = await axios.get<IResponseTokenParams>(
            `${TOKEN_LINK}?key=${responseObj.key}&oauth_token=${responseObj.oauth_token}&oauth_verifier=${responseObj.oauth_verifier}`
          );

          const responseData = response.data;
          localStorage.setItem("responseKey", responseData.Token || "");
          localStorage.setItem("firstName", responseData.FirstName || "");
          localStorage.setItem("lastName", responseData.LastName || "");
          console.log(responseData);
          navigate("/home"); // Używamy hooka navigate bezpośrednio
        } catch (error) {
          console.error("Error fetching data:", error);
        }
      };

      console.log(responseObj);
      getData();
    }
  }, [navigate]); // Dodajemy navigate do zależności useEffect

  return <div>Oczekiwanie na pobranie danych</div>;
};

export default LoadingPage;
