import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent implements OnInit {

  username:     string;
  password:     string;
  confPassword: string;
  email:        string;
  info:         string;
  url;s

  signup(): void {
    console.log("im alive");
    console.log(this.password === this.confPassword);
    console.log(this.username, this.email, this.info);
  }

  onSelectFile(event) {
    if (event.target.files && event.target.files[0]) {
      
      var reader = new FileReader();

      reader.readAsDataURL(event.target.files[0]);

      reader.onload = (event) => {
        this.url = event.target.result;
      }
    }
  }

  constructor() { }

  ngOnInit() {
  }

}
