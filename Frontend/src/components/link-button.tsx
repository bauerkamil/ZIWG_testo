import { useEffect } from "react";
import { Button } from "./ui/button";

interface IResponseTypeLogin {
  key?: string;
  oauth_token?: string;
  oauth_verifier?: string;
}

const LinkButton = () => {
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
      console.log(responseObj);
    }
  }, []);

  const handleClick = () => {
    window.open(
      "https://func-pwr-testo-dev.azurewebsites.net/api/auth?callback_url=http://localhost:5173"
    );
  };

  return <Button onClick={handleClick}>Zaloguj poprzez USOS</Button>;
};

export default LinkButton;
