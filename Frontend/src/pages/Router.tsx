import {createBrowserRouter} from "react-router-dom";
import Home from "@/pages/home/Home.tsx";
import ErrorPage from "@/pages/error-page/ErrorPage.tsx";

const Router = createBrowserRouter([
    {
        path: "/",
        element: <Home />,
        errorElement: <ErrorPage />,
        children: [],
    },
]);

export default Router;