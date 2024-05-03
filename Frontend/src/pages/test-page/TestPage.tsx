import Client from "@/api/Client";
import Navbar from "@/components/navbar/Navbar";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { NavbarPages } from "@/shared/enums/NavbarPages";
import { useEffect } from "react";
import { useParams } from "react-router-dom";

const TestPage: React.FC = () => {
    const { id } = useParams<{ id: string }>();

    useEffect(() => {
        const fetchData = async () => {
            if (id) {
                try {
                    const testData = await Client.getTest(id);
                    console.log(testData)
                } catch (error) {
                    console.error("An error occurred while fetching tests:", error);
                }
            }
        };

        fetchData();
    }, []);

    return (
        <div className="flex min-h-screen w-full flex-col">
            <Navbar page={NavbarPages.Home} />
            <main className="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
                <Card>
                    <CardHeader>
                        <CardTitle>Card Title</CardTitle>
                        <CardDescription>Card Description</CardDescription>
                    </CardHeader>
                    <CardContent>
                        <p>Card Content</p>
                    </CardContent>
                    <CardFooter>
                        <p>Card Footer</p>
                    </CardFooter>
                </Card>

            </main>
        </div>
    )
}

export default TestPage;