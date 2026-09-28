// Smoke tests for the Mwangaza app shell.
//
// These widgets build synchronously and do not touch the network or the
// shared_preferences / geolocator platform plugins.

import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mwangaza/auth/screens/login.dart';
import 'package:mwangaza/auth/screens/signup.dart';
import 'package:mwangaza/features/home/views/dashboard_view.dart';
import 'package:mwangaza/features/home/views/insight_view.dart';
import 'package:mwangaza/main.dart';

void main() {
  testWidgets('app boots into the sign up screen', (WidgetTester tester) async {
    await pumpScreen(tester, const MyApp());

    expect(find.byType(SignUpPage), findsOneWidget);
    expect(find.text('Sign up to access free weather data'), findsOneWidget);
    expect(find.text('Mwangaza'), findsOneWidget);
  });

  testWidgets('sign up screen links to the login screen', (
    WidgetTester tester,
  ) async {
    await pumpScreen(tester, const MyApp());

    final logInLink = find.text('Log In');
    await tester.ensureVisible(logInLink);
    await tester.pumpAndSettle();

    await tester.tap(logInLink);
    await tester.pumpAndSettle();

    expect(find.byType(LoginPage), findsOneWidget);
    expect(find.byType(SignUpPage), findsNothing);
  });

  testWidgets('dashboard renders its summary cards', (
    WidgetTester tester,
  ) async {
    await pumpScreen(tester, const MaterialApp(home: DashboardView()));

    expect(find.text('Dashboard'), findsOneWidget);
    expect(find.text('Total Users'), findsOneWidget);
    expect(find.text('Revenue'), findsOneWidget);
    expect(find.text('Orders'), findsOneWidget);
  });

  testWidgets('insights view renders its overview cards', (
    WidgetTester tester,
  ) async {
    await pumpScreen(tester, const MaterialApp(home: InsightView()));

    expect(find.text('Insights'), findsOneWidget);
    expect(find.text('Total Sessions'), findsOneWidget);
    expect(find.text('Conversion Rate'), findsOneWidget);
  });
}

/// Pumps [app] on a desktop-sized surface: the auth screens split the viewport
/// into two panes, which is too narrow for the default 800x600 test surface.
Future<void> pumpScreen(WidgetTester tester, Widget app) async {
  tester.view.physicalSize = const Size(1400, 1000);
  tester.view.devicePixelRatio = 1;
  addTearDown(tester.view.reset);

  await tester.pumpWidget(app);
}
