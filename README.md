#Rock Paper Scissors Kata
```
Rock Paper Scissors is a game involving two players making pre-defined hand gestures at each other. The gesture that each player uses is played against the other, with a winner being decided based on the rules being used.

The three gestures used in base Rock Paper Scissors are... well... rock, paper, and scissors. The way these are scored is as such: Rock beats Scissors, Scissors beats Paper, Paper beats Rock. It gets a lot more complicated when you introduce new gestures, but let's keep it simple for now.

As always, we want you to create a backend for the game that we can use to hook up to our many game clients we're going to be creating. Once again, feel free to use any front-end to test your program.

#Scenarios
- Rock Beats Scissors
- Scissors Beats Paper
- Paper Beats Rock
- Same Moves Result in Draw


```
# Function And Parameter Naming Conventions

## Directory Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
insurance
```

## File Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
api.go
```

## Package Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
insurance
```

## Test Function Name
- ใช้รูปแบบการตั้งชื่อฟังก์ชันเป็นแบบ **Snake_Case** เช่น
```
Test_RequestToOtherWebservice_Input_ApplicationForm_Should_Be_ApplicationResult
```

## Variable Name
- ชื่อตัวแปรเป็น **camelCase** เช่น
```
policy, applicationForms, applicationResult, policyResult
```

- ชื่อตัวแปรเก็บค่าที่เป็นพหูพจน์ ให้เติม s ต่อท้ายตัวแปรเสมอ เช่น
```
insurers
```

- ชื่อตัวแปร struct ให้ตั้งชื่อขึ้นต้นคำแรกด้วยตัวอักษรพิมพ์ใหญ่ ในรูปแบบ **camelCase** เช่น
```
ResultDay, ResultMonth
```

- ชื่อตัวแปร Constant ให้ตังชื่อเป็นตัวพิมพ์เล็กก่อน เว้นแต่เมื่อมีการใช้ข้าม package ถึงจะใช้ Capital Case เช่น
```
Hour, Minute, url
```

## ข้อตกลง Commit Message ร่วมกัน
`[Created]: สร้างไฟล์ใหม่`

`[Edited]: แก้ไข code ในไฟล์เดิมที่มีอยู่แล้ว รวมถึงกรณี refactor code`

`[Added]: กรณีเพิ่ม function, function test ใหม่เข้ามา`

`[Deleted]: ลบไฟล์ออก`

* ให้เขียนรายละเอียดด้วยว่าแก้ไขอะไรและทำที่ตรงไหน

## คำสั่ง Run Test
### ค่าสั่ง Run Acceptance Test (Robot)

```
robot duration.robot
```

### คำสั่ง Run Acceptance Test (API)
```
newman run filename
```

## Additional
[See more git and go command](https://github.com/ImKK-000/git-and-go-step)