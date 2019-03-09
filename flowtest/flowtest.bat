rem cmd /k cd/d "d:\gowork\src\github.com\3xxx\flowtest"
@echo off
d:
cd d:\gowork\src\github.com\3xxx\flowtest
bee run -gendoc=true -downdoc=true
cmd /k