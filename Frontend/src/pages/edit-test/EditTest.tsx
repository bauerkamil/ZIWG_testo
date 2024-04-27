import Combobox from "@/components/combobox/Combobox";
import Navbar from "@/components/navbar/Navbar";
import { Button, Input, Label } from "@/components/ui";
import { ICourse } from "@/shared/interfaces";
import { Plus } from "lucide-react";
import React from "react";

const EditTest: React.FC = () => {
  const [courses, setCourses] = React.useState<ICourse[]>([]);
  const [selectedCourse, setSelectedCourse] = React.useState<ICourse | null>(
    null
  );

  React.useEffect(() => {
    const mock: ICourse[] = [
      {
        id: "1",
        name: "Analiza matematyczna",
        teacher: "prof. Jan Kowalski",
      },
      {
        id: "2",
        name: "Algebra liniowa",
        teacher: "dr inż. Anna Nowak",
      },
    ];
    setCourses(mock);
    // setSelectedCourse(mock[0]);
  }, []);

  return (
    <div className="flex min-h-screen w-full flex-col">
      <Navbar />
      <main className="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
        <div className="grid gap-2 text-center">
          <div className="text-4xl">Edytuj testownik</div>
          {/* <p className="text-balance text-muted-foreground">Zacznij</p> */}
        </div>
        <div>
          <Label>Nazwa testownika</Label>
          <Input />
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
          <Input readOnly value={selectedCourse?.teacher} />
        </div>
        <Button className="gap-2">
          <Plus className="h-5 w-5" /> Dodaj nowe pytanie
        </Button>
      </main>
    </div>
  );
};

export default EditTest;
