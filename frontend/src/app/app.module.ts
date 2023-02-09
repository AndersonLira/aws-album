import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { FolderListComponent } from './folder-list/folder-list.component';
import { ImageCardComponent } from './image-card/image-card.component';

@NgModule({
  declarations: [
    AppComponent,
    FolderListComponent,
    ImageCardComponent
  ],
  imports: [
    BrowserModule, HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
