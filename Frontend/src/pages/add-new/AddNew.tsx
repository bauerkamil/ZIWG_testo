import Navbar from '@/components/navbar/Navbar';
import { Button } from '@/components/ui/button';
import { NavbarPages } from '@/shared/enums';
import React from 'react';

const AddNew: React.FC = () => {

  return (
  <div className="flex min-h-screen w-full flex-col">
    <Navbar page={NavbarPages.AddNew}/>
    <div className="flex items-center justify-center py-12">
        <div className="mx-auto grid w-[550px] gap-6">
          <div className="grid gap-2 text-center">
            <h1 className="text-6xl font-bold">DODAJ NOWE PYTANIA</h1>
            <p className="text-balance text-muted-foreground">
              Pozwól sobie i innym przetestować swoje umiejętności
            </p>
          </div>
          <div className="grid gap-4 mt-4">
            <Button >Dodaj z plików</Button>
            <Button >Dodaj ze zdjęć (OCR)</Button>
            <Button >Dodaj ręcznie</Button>

          </div>
        </div>
      </div>
  </div>
  );
}

export default AddNew;