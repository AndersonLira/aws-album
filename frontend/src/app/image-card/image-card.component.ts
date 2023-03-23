import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-image-card',
  templateUrl: './image-card.component.html',
  styleUrls: ['./image-card.component.css']
})
export class ImageCardComponent implements OnInit {
  @Input()
  image: string = ''
  selected = false;
  isImage = true;

  ngOnInit(): void {
    this.isImage = this.image.indexOf('.mp4?') == -1;
  }



  click() {
    this.selected = !this.selected;
  }
}
