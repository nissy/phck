Summary: PHCK is system health checker web server.
Name: phck
Version: 0.2
Release: 1%{?dist}
License: GNU General Public License, Version 3
Group: System Environment/Base
URL: https://github.com/ngc224/phck

Source0: https://github.com/ngc224/phck/releases/download/%{version}/phck-%{version}_linux_amd64.tar.gz
Source1: phck.conf
Source2: phck.service

BuildRoot: %{_tmppath}/%{name}-root

%description
PHCK is system health checker web server.

%prep
%setup -q

%install
rm -rf $RPM_BUILD_ROOT

mkdir -p $RPM_BUILD_ROOT%{_sbindir}
install -m 755 phck $RPM_BUILD_ROOT%{_sbindir}

mkdir -p $RPM_BUILD_ROOT%{_sysconfdir}/phck
install -m 644 %SOURCE1 $RPM_BUILD_ROOT%{_sysconfdir}/phck/phck.conf

mkdir -p $RPM_BUILD_ROOT%{_unitdir}
install -m 644 %SOURCE2 $RPM_BUILD_ROOT%{_unitdir}/phck.service

%clean
rm -rf $RPM_BUILD_ROOT

%files
%defattr(-,root,root)
%{_sbindir}/phck
%config(noreplace) %{_sysconfdir}/phck/phck.conf
%{_unitdir}/phck.service

%changelog
* Tue Jun 7 2016 Yoshihiko Nishida <nishida@ngc224.org> - 7-2.el7.centos
- Build for CentOS-7.2
