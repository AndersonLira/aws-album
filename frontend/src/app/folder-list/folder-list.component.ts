import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FilesService } from '../files.service';


@Component({
  selector: 'app-folder-list',
  templateUrl: './folder-list.component.html',
  styleUrls: ['./folder-list.component.css']
})

export class FolderListComponent implements OnInit{
  files: any
  items: any
  cache = {} as any;

  @Output()
  currentFiles = new EventEmitter<string[]>()
  currentPath = '';


  constructor(private readonly service: FilesService){}

  ngOnInit(): void {
    this.service.getFiles().subscribe((it: any) => {
      this.files = it;
      this.prepareItems(this.files)
    })
  }

  click(event:any,name:any) {
    let selected = this.items.filter((it:any) => it.name == name)[0];
    this.prepareItems(selected);
    this.items = this.items.map((it:any) => { return {...it,selected: it.name == name}});
  }

  prepareItems(current:any) {
    const list = current.name ? current.files : current;
    if(current.name) {
      this.currentPath = this.currentPath + current.name + "/";
    }
    this.items = [];
    const currentFiles = [] as string[];
    list.forEach((element: any) => {
      if(!element.isFile){
        this.items.push(element)
      }else{
        currentFiles.push(this.currentPath +  element.name);
      }
    });
    this.getUrls(currentFiles);
  }

  private getUrls(currentFiles: any) {
    const key = this.currentPath + '-';
    const aux = this.cache[key];
    if(!aux){
      this.service.getUrls(currentFiles).subscribe((it : string[]) => {
        this.cache[key] = it;
        this.currentFiles.emit(it);
      })
    }else {
      this.currentFiles.emit(aux);
    }

  }

  back() {
    this.prepareItems(this.files);
    this.currentPath = '';
  }


}
