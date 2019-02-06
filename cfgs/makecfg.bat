@echo off
set makecmd=go run ..\cmd\makecfg\main.go

REM without-enum
%makecmd% -o gtk-server.cfg

%makecmd% -os=unix -lib=gtk1 -o gtk-server-unix-gtk1.cfg
%makecmd% -os=unix -lib=gtk2 -o gtk-server-unix-gtk2.cfg
%makecmd% -os=unix -lib=gtk3 -o gtk-server-unix-gtk3.cfg
%makecmd% -os=unix -lib=xforms -o gtk-server-unix-xforms.cfg
%makecmd% -os=unix -lib=motif -o gtk-server-unix-motif.cfg

%makecmd% -os=win -lib=gtk1 -o gtk-server-win-gtk1.cfg
%makecmd% -os=win -lib=gtk2 -o gtk-server-win-gtk2.cfg
%makecmd% -os=win -lib=gtk3 -o gtk-server-win-gtk3.cfg

%makecmd% -os=mac -lib=gtk2 -o gtk-server-mac-gtk2.cfg

REM with-enum
%makecmd%  -with-enum -o gtk-server-with-enum.cfg
%makecmd% -os=unix -lib=gtk1 -with-enum -o gtk-server-unix-gtk1-with-enum.cfg
%makecmd% -os=unix -lib=gtk2 -with-enum -o gtk-server-unix-gtk2-with-enum.cfg
%makecmd% -os=unix -lib=gtk3 -with-enum -o gtk-server-unix-gtk3-with-enum.cfg
%makecmd% -os=unix -lib=xforms -with-enum -o gtk-server-unix-xforms-with-enum.cfg
%makecmd% -os=unix -lib=motif -with-enum -o gtk-server-unix-motif-with-enum.cfg

%makecmd% -os=win -lib=gtk1 -with-enum -o gtk-server-win-gtk1-with-enum.cfg
%makecmd% -os=win -lib=gtk2 -with-enum -o gtk-server-win-gtk2-with-enum.cfg
%makecmd% -os=win -lib=gtk3 -with-enum -o gtk-server-win-gtk3-with-enum.cfg

%makecmd% -os=mac -lib=gtk2 -with-enum -o gtk-server-mac-gtk2-with-enum.cfg
