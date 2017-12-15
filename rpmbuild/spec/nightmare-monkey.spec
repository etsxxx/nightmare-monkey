%define _binaries_in_noarch_packages_terminate_build 0

Summary: Nightmare Monkey makes accidents to OS and processes.
Name:    nightmare-monkey
Version: 0.0.1
Release: 1
License: MIT
URL:     https://github.com/etsxxx/nightmare-monkey

%define INSTALLDIR %{buildroot}/usr/local/bin

%description
%{summary}

%prep

%build
cd /root/go/src/github.com/etsxxx/nightmare-monkey
make VERSION=%{version}

%install
%{__rm} -rf %{buildroot}
mkdir -p %{INSTALLDIR}
%{__install} -Dp -m0755 /root/go/src/github.com/etsxxx/nightmare-monkey/bin/* %{INSTALLDIR}

%clean
%{__rm} -rf %{buildroot}

%post

%files
%defattr(-,root,root)
/usr/local/bin
