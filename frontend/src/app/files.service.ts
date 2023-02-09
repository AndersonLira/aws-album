import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'

@Injectable({
  providedIn: 'root'
})
export class FilesService {

  constructor(private readonly http: HttpClient) { 

  }

  getFiles(): any {
    return this.http.get('http://localhost:7000/folders/')
  }

  getUrls(list: any): any {
    let aux = "";
    list.forEach((file:any) => {
      aux += "values=" + file + "&";
    })
    return this.http.get('http://localhost:7000/urls?' + aux)
  }
}
