import { Component } from '@angular/core';
import { FilesService } from './files.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'frontend';
  images: string[] = [];
  constructor(private readonly service: FilesService) {}

  currentFiles(event:any){
    this.service.getUrls(event).subscribe((it : string[]) => {
      this.images = it;
    })
  }
}
