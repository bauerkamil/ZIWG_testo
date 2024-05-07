import React from "react";
import Combobox from "@/components/combobox/Combobox";
import Navbar from "@/components/navbar/Navbar";
import { Button, Input, Label } from "@/components/ui";
import { ICourse, IQuestion, ITest } from "@/shared/interfaces";
import { Plus } from "lucide-react";
import EditQuestion from "./question/EditQuestion";
import { useParams } from "react-router-dom";
import Client from "@/api/Client";

const EditTest: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [questions, setQuestions] = React.useState<IQuestion[]>([]);
  const [courses, setCourses] = React.useState<ICourse[]>([]);
  const [test, setTest] = React.useState<ITest>();
  const [selectedCourse, setSelectedCourse] = React.useState<ICourse | null>(
    null
  );

  React.useEffect(() => {
    const fetchData = async () => {
      if (id) {
        try {
          const testData = await Client.getTest(id);
          const coursesData = await Client.getCourses();
          console.log(testData);
          console.log(coursesData);
          setTest(testData);
          setQuestions(testData.questions ?? []);
          setCourses(coursesData)
        } catch (error) {
          console.error("An error occurred while fetching tests:", error);
        }
      }
    };

    fetchData();
  }, [id]);

  return (
    <div className="flex min-h-screen w-full flex-col">
      <Navbar />
      <main className="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
        <div className="grid gap-2 text-center">
          <div className="text-4xl">Edytuj testownik</div>
        </div>
        <div>
          <Label>Nazwa testownika</Label>
          <Input defaultValue={test?.name}/>
        </div>
        <div>
          <Label>Wybierz kurs</Label>
          <Combobox
            items={courses}
            selectedItem={selectedCourse}
            onItemSelected={setSelectedCourse}
            keyPath="id"
            valuePath="name"
          />
        </div>
        <div>
          <Label className="text-muted-foreground">Prowadzący</Label>
          <Input
            readOnly
            value={selectedCourse?.teacher ? selectedCourse.teacher.name + selectedCourse.teacher.surname : ""}
            placeholder="Wybierz kurs żeby zobaczyć prowadzącego"
          />
        </div>
        <div>
          <Label>Pytania</Label>
          <div className="grid gap-4">
            {questions.map((question, index) => (
              <EditQuestion
                key={question.id}
                index={index}
                question={question}
              />
            ))}
          </div>
        </div>
        <Button className="gap-2">
          <Plus className="h-5 w-5" /> Dodaj nowe pytanie
        </Button>
      </main>
    </div>
  );
};

export default EditTest;
