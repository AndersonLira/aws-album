import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FilesService } from '../files.service';

@Component({
  selector: 'app-folder-list',
  templateUrl: './folder-list.component.html',
  styleUrls: ['./folder-list.component.css']
})
export class FolderListComponent implements OnInit{
  items: any

  @Output()
  currentFiles = new EventEmitter<string[]>()
  private currentPath = "";

  constructor(private readonly service: FilesService){}

  ngOnInit(): void {
    this.items = [];
    const currentFiles = [] as string[];
    this.service.getFiles().subscribe((it: any) => {
      it.forEach((element: any) => {
        if(!element.isFile){
          this.items.push({name: element.name})
        }else{
          currentFiles.push(this.currentPath +  element.name);
        }
      });
      this.currentFiles.emit(currentFiles);
    })
  }

}
