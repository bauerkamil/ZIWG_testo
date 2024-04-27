import React from "react";

import { useAuth } from "@/shared/hooks/auth/useAuth";
import Navbar from "@/components/navbar/Navbar";
import { NavbarPages } from "@/shared/enums";
import { ITestInfo } from "@/shared/interfaces";
import TestCard from "./test-card/TestCard";

const Home: React.FC = () => {
  const { user } = useAuth();

  const tests: ITestInfo[] = [
    {
      id: "1",
      name: "Kolokwium 1",
      course: "Analiza matematyczna",
      teacher: "prof. Jan Kowalski",
      lastModified: new Date(),
    },
    {
      id: "2",
      name: "Kolokwium 2",
      course: "Algebra liniowa",
      teacher: "dr inż. Anna Nowak",
      lastModified: new Date(),
    },
  ];

  return (
    <div className="flex min-h-screen w-full flex-col">
      <Navbar page={NavbarPages.Home} />
      <main className="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
        <div className="grid gap-2 text-center">
          <div className="text-4xl">
            Cześć, {user?.firstName} {user?.lastName}
          </div>
          <p className="text-balance text-muted-foreground">
            Zacznij się uczyć (najlepiej tak z dzień przed kolosem)
          </p>
        </div>
        <div className="text-2xl">Dostępne testowniki na ten semestr</div>
        <div className="grid gap-4 md:grid-cols-2">
          {tests.map((test) => (
            <TestCard key={test.id} test={test} />
          ))}
        </div>
      </main>
    </div>
  );
};

export default Home;
